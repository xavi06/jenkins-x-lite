package jenkins

import (
	"fmt"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/dockerfile"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
)

// JAuth struct
type JAuth struct {
	URL  string
	User string
	Pass string
}

func getJenkinsHost() string {
	jenkinsHost := os.Getenv("JENKINS_URL")
	if jenkinsHost == "" {
		jenkinsHost = "http://localhost:8080"
	}
	return jenkinsHost
}

// GetJAuth func
func GetJAuth() *JAuth {
	url := os.Getenv("JENKINS_URL")
	if url == "" {
		url = "http://localhost:8080"
	}
	user := os.Getenv("JENKINS_USER")
	if user == "" {
		user = "admin"
	}
	pass := os.Getenv("JENKINS_PASS")
	if pass == "" {
		pass = "admin"
	}
	return &JAuth{URL: url, User: user, Pass: pass}
}

func getStep(name string, config *util.TomlConfig) (step string) {
	switch name {
	case "git":
		var gitSecret = "1dc99102-60f4-4cee-bc86-cbcdf6c21071"
		if config.Jenkins.GitSecret != "" {
			gitSecret = config.Jenkins.GitSecret
		}
		step += "        stage('git clone') {\n"
		step += "            steps {\n"
		step += fmt.Sprintf("                git credentialsId: '%s', url: '%s', branch: '%s'\n", gitSecret, config.Git.Repo, config.Git.Branch)
		step += "            }\n"
		step += "        }\n"
	case "prebuild":
		step += "        stage('run prebuild') {\n"
		step += "            steps {\n"
		step += fmt.Sprintf("                sh \"sh /var/lib/jenkins/dockerfiles/%s/prebuild.sh %s\"\n", config.Chart.Name, config.Chart.Name)
		step += "            }\n"
		step += "        }\n"
	case "mvn-package":
		step += "        stage('mvn package') {\n"
		step += "            steps {\n"
		step += "                sh \"mvn clean package -U\"\n"
		step += "            }\n"
		step += "        }\n"
	case "mvn-package-skip":
		step += "        stage('mvn package') {\n"
		step += "            steps {\n"
		step += "                sh \"mvn clean package -Dmaven.test.skip=true -U\"\n"
		step += "            }\n"
		step += "        }\n"
	case "get-dockerfile":
		dockerfileURL := dockerfile.GetFileServer()
		step += "        stage('get dockerfile') {\n"
		step += "            steps {\n"
		step += fmt.Sprintf("                sh \"curl -o Dockerfile %s/files/dockerfiles/%s\"\n", dockerfileURL, config.Chart.Name)
		step += "            }\n"
		step += "        }\n"
	case "build-and-push":
		var dockerSecret = "cc41a11b-a2f4-4093-97cf-92017b68977c"
		if config.Jenkins.DockerSecret != "" {
			dockerSecret = config.Jenkins.DockerSecret
		}
		step += "        stage('docker build and push') {\n"
		step += "            steps {\n"
		step += "                script{\n"
		step += fmt.Sprintf("                    docker.withRegistry(\"http://%s\", '%s') {\n", config.Helm.Image.Repository, dockerSecret)
		step += fmt.Sprintf("                        def customImage = docker.build(\"%s:${params.tag}\")\n", config.Helm.Image.Name)
		step += "                        customImage.push()\n"
		step += "                    }\n"
		step += "                }\n"
		step += "            }\n"
		step += "        }\n"
	case "save-image-tag":
		step += "        stage('save image tag') {\n"
		step += "            steps {\n"
		step += fmt.Sprintf("                sh \"ETCDCTL_API=3 etcdctl --endpoints=[10.30.65.71:4001] put /dockerimage/%s '${params.tag}'\"\n", config.Helm.Image.Name)
		step += "            }\n"
		step += "        }\n"
	case "helm":
		step += "        stage('helm deploy') {\n"
		step += "            steps {\n"
		step += "                sh \"sudo /usr/local/bin/helm repo update\"\n"
		step += "                script{\n"
		step += "					if (params.action == 'install') {\n"
		step += fmt.Sprintf("						sh \"sudo /usr/local/bin/helm --kubeconfig /root/.kube/config_%s install --name %s --namespace %s --set image.tag='${params.tag}' %s\"\n", config.Cluster, config.Chart.Name, config.Namespace, config.Helmrepo+"/"+config.Chart.Name)
		step += "					} else {\n"
		step += fmt.Sprintf("						sh \"sudo /usr/local/bin/helm --kubeconfig /root/.kube/config_%s upgrade %s --set image.tag='${params.tag}' %s\"\n", config.Cluster, config.Chart.Name, config.Helmrepo+"/"+config.Chart.Name)
		step += "					}\n"
		step += "                }\n"
		step += "            }\n"
		step += "        }\n"
	case "rsync":
		step += "        stage('rsync data') {\n"
		step += "            steps {\n"
		for _, syncAddr := range config.Jenkins.SyncAddr {
			//step += fmt.Sprintf("                sh 'sudo rsync -arvt --delete --exclude=.git --exclude=.svn --exclude=/Jenkinsfile --exclude=README.md --exclude=/WEB-INF/ ./src/main/webapp/ root@%s:/data/NASShare/www/%s/'\n", syncAddr, config.Chart.Name)
			step += fmt.Sprintf("                sh 'sudo /usr/bin/ansible %s -m synchronize -a \"src=./src/main/webapp/ dest=/data/NASShare/www/%s/ delete=yes rsync_opts=--exclude=.git,--exclude=.svn,--exclude=/Jenkinsfile,--exclude=README.md,--exclude=/WEB-INF/\"'\n", syncAddr, config.Chart.Name)
		}
		step += "            }\n"
		step += "        }\n"
	}

	return step
}

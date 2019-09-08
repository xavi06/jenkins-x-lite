package tekton

import (
	"fmt"

	"github.com/xavi06/jenkins-x-lite/pkg/dockerfile"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	yaml "gopkg.in/yaml.v2"
)

const (
	apiVersion     = "tekton.dev/v1alpha1"
	serviceAccount = "tekton-git-and-docker-admin"
	namespace      = "tekton-pipelines"
)

// CreatePipelineResourceString func
func CreatePipelineResourceString(name, ns, stype string, params []PipelineResourceParam) (string, error) {
	var metaData *MetaData
	var labels map[string]string
	labels = make(map[string]string)
	labels["app"] = name
	metaData = &MetaData{
		Name:      name + "-" + stype,
		Namespace: ns,
		Labels:    labels,
	}

	var spec *PipelineResourceSpec
	spec = &PipelineResourceSpec{
		Type:   stype,
		Params: params,
	}
	var pipelineResource *PipelineResource
	pipelineResource = &PipelineResource{
		APIVersion: apiVersion,
		Kind:       "PipelineResource",
		MetaData:   *metaData,
		Spec:       *spec,
	}
	yamlconfig, err := yaml.Marshal(&pipelineResource)
	if err != nil {
		return "", err
	}
	return string(yamlconfig), nil
}

// GetSteps func
func getStep(name, tag string, config *util.TomlConfig) (step TaskSpecStep) {
	switch name {
	case "get-dockerfile":
		dockerfileURL := dockerfile.GetFileServer()
		getdockerfileCmd := fmt.Sprintf("wget -O Dockerfile %s/files/dockerfiles/%s", dockerfileURL, config.Chart.Name)
		step = TaskSpecStep{
			Name:       name,
			Image:      "alpine:3.8",
			WorkingDir: "/workspace/app",
			Command:    []string{"/bin/sh"},
			Args:       []string{"-c", getdockerfileCmd},
		}
	case "mvn-package":
		step = TaskSpecStep{
			Name: name,
			//Image:      "registry.cn-hangzhou.aliyuncs.com/zjops/tekton-maven-executor:3.5.4-0",
			Image:      config.Helm.Image.Repository + "/zj/tekton/maven-executor:3.5.4-0",
			WorkingDir: "/workspace/app",
			Command:    []string{"mvn"},
			Args:       []string{"clean", "package"},
		}
	case "build-and-push":
		step = TaskSpecStep{
			Name:  name,
			Image: "gcr.io/kaniko-project/executor",
			Env: []TaskSpecStepsEnv{
				{
					Name:  "DOCKER_CONFIG",
					Value: "/builder/home/.docker/",
				},
			},
			Command: []string{"/kaniko/executor"},
			Args: []string{
				"--dockerfile=${inputs.params.pathToDockerFile}",
				"--destination=${outputs.resources.builtImage.url}",
				"--context=${inputs.params.pathToContext}",
				"--skip-tls-verify",
			},
		}
	case "helm-install":
		helmCmd := fmt.Sprintf("helm repo update && helm install --name %s --namespace %s --set image.tag=%s chartmuseum/%s", config.Chart.Name, config.Namespace, tag, config.Chart.Name)
		//fmt.Println(helmCmd)
		step = TaskSpecStep{
			Name: name,
			//Image: "registry.cn-hangzhou.aliyuncs.com/zjops/tekton-test-kubectl-helm-executor:1.0.0",
			Image: config.Helm.Image.Repository + "/zj/tekton/kubectl-helm-executor:1.0.0",
			Env: []TaskSpecStepsEnv{
				{
					Name:  "HELM_HOME",
					Value: "/root/.helm",
				},
				{
					Name:  "HELMREPO",
					Value: config.Helmrepo,
				},
			},
			Command: []string{"/bin/sh"},
			Args:    []string{"-c", helmCmd},
		}
	case "helm-upgrade":
		helmCmd := fmt.Sprintf("helm repo update && helm upgrade %s --set image.tag=%s chartmuseum/%s", config.Chart.Name, tag, config.Chart.Name)
		step = TaskSpecStep{
			Name: name,
			//Image: "registry.cn-hangzhou.aliyuncs.com/zjops/tekton-test-kubectl-helm-executor:1.0.0",
			Image: config.Helm.Image.Repository + "/zj/tekton/kubectl-helm-executor:1.0.0",

			Env: []TaskSpecStepsEnv{
				{
					Name:  "HELM_HOME",
					Value: "/root/.helm",
				},
				{
					Name:  "HELMREPO",
					Value: config.Helmrepo,
				},
			},
			Command: []string{"/bin/sh"},
			Args:    []string{"-c", helmCmd},
		}
	case "rsync":
		syncCmd := fmt.Sprintf("rsync -arvt --delete --exclude=.git --exclude=.svn --exclude=/Jenkinsfile --exclude=README.md --exclude=/WEB-INF/ ./src/main/webapp/ root@172.16.162.8:/data/NASShare/www/%s/", config.Chart.Name)
		//fmt.Println(syncCmd)
		step = TaskSpecStep{
			Name:       name,
			Image:      config.Helm.Image.Repository + "/zj/tekton/rsync-ssh:1.0.0",
			WorkingDir: "/workspace/app",
			Command:    []string{"/bin/sh"},
			Args:       []string{"-c", syncCmd},
		}
	default:
		return

	}
	return
}

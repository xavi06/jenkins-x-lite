package jenkins

import (
	"errors"
	"strings"

	"github.com/xavi06/jenkins-x-lite/pkg/util"
)

// CreateMavenConfig func
func CreateMavenConfig(gitrepo, gitbranch, dockerregistry, imagename string) string {
	return `<?xml version='1.1' encoding='UTF-8'?>
	<flow-definition plugin="workflow-job@2.29">
	  <actions>
		<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.3.2"/>
		<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.3.2">
		  <jobProperties/>
		  <triggers/>
		  <parameters/>
		  <options/>
		</org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
	  </actions>
	  <description></description>
	  <keepDependencies>false</keepDependencies>
	  <properties>
		<jenkins.model.BuildDiscarderProperty>
		  <strategy class="hudson.tasks.LogRotator">
			<daysToKeep>-1</daysToKeep>
			<numToKeep>20</numToKeep>
			<artifactDaysToKeep>-1</artifactDaysToKeep>
			<artifactNumToKeep>-1</artifactNumToKeep>
		  </strategy>
		</jenkins.model.BuildDiscarderProperty>
	  </properties>
	  <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.62">
		<script>pipeline {
		agent any
		environment {
			PATH = &quot;/usr/local/maven/bin:/usr/local/java/jdk1.8.0_171/bin/:$PATH&quot;
		}
		stages {
			stage(&apos;git clone&apos;) {
				steps {
					git credentialsId: &apos;1dc99102-60f4-4cee-bc86-cbcdf6c21071&apos;, url: &apos;` + gitrepo + `&apos;, branch: &apos;` + gitbranch + `&apos;
				}
			}
	
			stage(&apos;mvn package&apos;) {
				steps {
					sh &quot;mvn package&quot;
				}
			}
	
			stage(&apos;docker build&amp;push&apos;){
				steps {
					script{
						docker.withRegistry(&quot;http://` + dockerregistry + `&quot;, &apos;cc41a11b-a2f4-4093-97cf-92017b68977c&apos;) {
	
							def customImage = docker.build(&quot;` + imagename + `:auto-${currentBuild.startTimeInMillis}&quot;)
	
							/* Push the container to the custom Registry */
							customImage.push()
						}
					}
				}
			}
	
			stage(&apos;kubernetes deploy&apos;) {
				steps {
					sh &quot;sudo rm -rf ./target&quot;
				}
			}
	
		}
	}
	</script>
		<sandbox>true</sandbox>
	  </definition>
	  <triggers/>
	  <disabled>false</disabled>
	</flow-definition>`
}

// CreateMavenJenkinsfile func
func CreateMavenJenkinsfile(gitrepo, gitbranch, dockerregistry, imagename, helmrelease, helmchart, namespace, dockerfileurl, cluster string) error {

	if util.CheckFileExsit("Jenkinsfile") {
		return errors.New("file exist! Delete it first")
	}

	jenkinsfile := `pipeline {
  agent any
	environment {
		PATH = "/usr/local/bin:/usr/local/maven/bin:/usr/local/java/jdk1.8.0_171/bin/:$PATH"
	}

	stages {
		stage('git clone') {
			steps {
				git credentialsId: '1dc99102-60f4-4cee-bc86-cbcdf6c21071', url: '` + gitrepo + `', branch: '` + gitbranch + `'
		  }
	  }

		//stage('run preinstall') {
			//steps {
				//sh "sh /var/lib/jenkins/dockerfiles/` + helmrelease + `/prebuild.sh ` + helmrelease + `"
			//}
		//}
	  
		stage('mvn package') {
			steps {
				sh "mvn clean package -U"
			}
		}
	
		stage('download Dockerfile') {
			steps {
				sh "curl -o Dockerfile ` + dockerfileurl + `/files/dockerfiles/` + helmrelease + `"
			}
		}

		stage('docker build and push'){
			steps {
				script{
					docker.withRegistry("http://` + dockerregistry + `", 'cc41a11b-a2f4-4093-97cf-92017b68977c') {

						def customImage = docker.build("` + imagename + `:${params.tag}")

						/* Push the container to the custom Registry */
						customImage.push()
					}
				}
			}
		}

		//stage('save image tag'){
            //steps {
                //sh "ETCDCTL_API=3 etcdctl --endpoints=[10.30.65.71:4001] put /dockerimage/` + imagename + ` '${params.tag}'"
            //}
		//}
		
		stage('kubernetes deploy') {
			steps {
				sh "sudo /usr/local/bin/helm repo update"
				script {
					if (params.action == 'install') {
						sh "sudo /usr/local/bin/helm --kubeconfig /root/.kube/config_` + cluster + ` install --name ` + helmrelease + ` --namespace ` + namespace + ` --set image.tag='${params.tag}' ` + helmchart + `"
					} else {	
						sh "sudo /usr/local/bin/helm --kubeconfig /root/.kube/config_` + cluster + ` upgrade ` + helmrelease + ` --set image.tag='${params.tag}' ` + helmchart + `"
					}
				}
			}
		}

		//stage('rsync data') {
			//steps {
			    //sh 'sudo rsync -arvt --delete --exclude=.git --exclude=.svn --exclude=/Jenkinsfile --exclude=README.md --exclude=/WEB-INF/ ./src/main/webapp/ root@10.30.70.59:/data/NASShare/www/` + helmrelease + `/'
			//}
		//}
	}
}`

	return util.WriteFileString("Jenkinsfile", jenkinsfile)
}

// CreatePipeline func
func CreatePipeline(jenkinsfile string) string {
	// Replace " and '
	jenkinsfile = strings.Replace(jenkinsfile, "\"", "&quot;", -1)
	jenkinsfile = strings.Replace(jenkinsfile, "'", "&apos;", -1)

	return `<?xml version='1.1' encoding='UTF-8'?>
	<flow-definition plugin="workflow-job@2.29">
	  <actions>
			<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.3.2"/>
			<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.3.2">
			  <jobProperties/>
			  <triggers/>
			  <parameters/>
			  <options/>
			</org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
	  </actions>
	  <description></description>
	  <keepDependencies>false</keepDependencies>
	  <properties>
	  <hudson.model.ParametersDefinitionProperty>
      <parameterDefinitions>
        <hudson.model.StringParameterDefinition>
          <name>action</name>
          <description></description>
          <defaultValue>update</defaultValue>
          <trim>false</trim>
		</hudson.model.StringParameterDefinition>
		<hudson.model.StringParameterDefinition>
		  <name>tag</name>
		  <description></description>
		  <defaultValue>latest</defaultValue>
		  <trim>false</trim>
	    </hudson.model.StringParameterDefinition>
      </parameterDefinitions>
		</hudson.model.ParametersDefinitionProperty>
		<hudson.plugins.jira.JiraProjectProperty plugin="jira@3.0.7"/>
    <hudson.security.AuthorizationMatrixProperty>
      <inheritanceStrategy class="org.jenkinsci.plugins.matrixauth.inheritance.InheritParentStrategy"/>
      <permission>hudson.model.Item.Build:authenticated</permission>
      <permission>hudson.model.Item.Cancel:authenticated</permission>
      <permission>hudson.model.Item.Read:authenticated</permission>
      <permission>hudson.model.Item.Workspace:authenticated</permission>
      <permission>hudson.scm.SCM.Tag:authenticated</permission>
    </hudson.security.AuthorizationMatrixProperty>
			<jenkins.model.BuildDiscarderProperty>
			  <strategy class="hudson.tasks.LogRotator">
					<daysToKeep>-1</daysToKeep>
					<numToKeep>20</numToKeep>
					<artifactDaysToKeep>-1</artifactDaysToKeep>
					<artifactNumToKeep>-1</artifactNumToKeep>
			  </strategy>
			</jenkins.model.BuildDiscarderProperty>
	  </properties>
	  <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.62">
			<script>` + jenkinsfile + `
	</script>
			<sandbox>true</sandbox>
	  </definition>
	  <triggers/>
	  <disabled>false</disabled>
	</flow-definition>`
}

// CreateJenkinsfile func
func CreateJenkinsfile(config *util.TomlConfig) error {
	if util.CheckFileExsit("Jenkinsfile") {
		return errors.New("file exist! Delete it first")
	}
	var jenkinsfile string
	jenkinsfile += "pipeline {\n"
	jenkinsfile += "    agent any\n"
	jenkinsfile += "    environment {\n"
	jenkinsfile += "        PATH = \"/usr/local/bin:/usr/local/maven/bin:/usr/local/java/jdk1.8.0_171/bin/:$PATH\"\n"
	jenkinsfile += "    }\n"
	jenkinsfile += "    stages {\n"
	// check steps
	steps := config.Jenkins.Steps
	for _, step := range steps {
		stepInfo := getStep(step, config)
		if stepInfo == "" {
			continue
		}
		jenkinsfile += stepInfo
	}
	jenkinsfile += "    }\n"
	jenkinsfile += "}"

	return util.WriteFileString("Jenkinsfile", jenkinsfile)
}

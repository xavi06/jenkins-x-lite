package cmd

import (
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/chartmuseum"
	"github.com/xavi06/jenkins-x-lite/pkg/dockerfile"
	"github.com/xavi06/jenkins-x-lite/pkg/helm"
	"github.com/xavi06/jenkins-x-lite/pkg/jenkins"
	"github.com/xavi06/jenkins-x-lite/pkg/tekton"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

const (
	configFile = "app.toml"
)

// Init func
// jxl init
func Init(c *cli.Context) error {
	log.Println("Parse toml config")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Init Helm Chart
	log.Println("Init Helm Chart")
	chartmuseumURL := chartmuseum.GetChartmuseumURL()
	err = helm.InitCharts("helloworld-maven", config.Chart.Name, chartmuseumURL+"/charts/"+config.Appbase)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Update Helm values.yaml")
	_, err = helm.ValuesConfig(config.Helm, "charts/"+config.Chart.Name+"/values.yaml")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	//fmt.Println(cf)

	log.Println("Update Helm Chart.yaml")
	_, err = helm.ChartsConfig(config.Chart, "charts/"+config.Chart.Name+"/Chart.yaml")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Check pipeline
	if config.Pipeline == "jenkins" {
		log.Println("Init Jenkinsfile")
		//dockerfileurl := dockerfile.GetFileServer()
		err = jenkins.CreateJenkinsfile(config)
		//err = jenkins.CreateMavenJenkinsfile(config.Git.Repo, config.Git.Branch, config.Helm.Image.Repository, config.Helm.Image.Name, config.Chart.Name, config.Helmrepo+"/"+config.Chart.Name, config.Namespace, dockerfileurl, config.Cluster)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	} else if config.Pipeline == "tekton" {
		log.Println("Init tekton pipeline")
		err = tekton.CreateResourceGit(config)
		if err != nil {
			log.Fatalln(err)
		}
		tag, err := tekton.CreateResourceImage(config)
		if err != nil {
			log.Fatalln(err)
		}
		err = tekton.CreateTask(config, "install", tag)
		if err != nil {
			log.Fatalln(err)
		}
		err = tekton.CreateTaskRun(config)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("Init Dockerfile")
	dockerf := dockerfile.CreateDockerfileString(config.Docker)
	err = dockerfile.CreateDockerfile(dockerf)
	if err != nil {
		log.Fatalln(err)
	}
	// END OF INIT
	log.Println("App init success!")
	return nil
}

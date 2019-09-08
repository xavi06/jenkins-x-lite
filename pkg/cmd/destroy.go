package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/kubernetes"

	"github.com/xavi06/jenkins-x-lite/pkg/chartmuseum"
	"github.com/xavi06/jenkins-x-lite/pkg/helm"
	"github.com/xavi06/jenkins-x-lite/pkg/jenkins"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// Destroy func
func Destroy(c *cli.Context) error {
	var sure string
	sure = "N"
	fmt.Printf("Are you sure? Y/N (Default: N) ")
	fmt.Scanln(&sure)
	if sure != "Y" {
		fmt.Println("Cancelled")
		return nil
	}
	log.Println("Begin destroy app")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Delete local helm chart")
	os.RemoveAll("charts")

	log.Println("Delete Dockerfile")
	os.Remove("Dockerfile")

	log.Println("Delete remote helm chart(chartmuseum)")
	chartmuseum.DeleteCharts(config.Chart.Name, config.Chart.Version)
	// Check pipeline
	if config.Pipeline == "jenkins" {
		log.Println("Delete Jenkinsfile")
		os.Remove("Jenkinsfile")
		log.Println("Delete jenkins job")
		jauth := jenkins.GetJAuth()
		jks := jenkins.NewJenkins(jauth)
		err = jks.DeleteJob(config.Chart.Name)
		if err != nil {
			log.Fatalln(err)
		}
	} else if config.Pipeline == "tekton" {
		log.Println("Delete tekton pipeline")
		res, err := kubernetes.KubeOper("delete", "pipelines/")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(res)
		//
		os.RemoveAll("pipelines")
	}

	log.Println("Delete helm chart")
	hcli := helm.NewHCli()
	hcli.Conn()
	_, err = hcli.Delete(config.Chart.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

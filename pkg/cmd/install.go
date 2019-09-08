package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/xavi06/jenkins-x-lite/pkg/chartmuseum"
	"github.com/xavi06/jenkins-x-lite/pkg/dockerfile"
	"github.com/xavi06/jenkins-x-lite/pkg/jenkins"
	"github.com/xavi06/jenkins-x-lite/pkg/kubernetes"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// Install func
// jxl install
func Install(c *cli.Context) error {
	//chartmuseumURL := chartmuseum.GetChartmuseumURL()
	log.Println("Parse toml config")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// helm package (compress)
	log.Println("Helm package (compress helm chart)")
	chartTgz := fmt.Sprintf("charts/%s-%s.tgz", config.Chart.Name, config.Chart.Version)
	chartDir := fmt.Sprintf("charts/%s", config.Chart.Name)
	err = util.Compress(chartTgz, chartDir)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// upload chart to chartmuseum
	log.Println("Upload to chartmuseum")
	chartdata, err := util.ReadFileBytes(chartTgz)
	err = chartmuseum.UploadCharts(chartdata)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Upload Dockerfile to fileserver
	log.Println("Upload Dockerfile")
	err = dockerfile.UploadDockerfile(config.Chart.Name)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Check pipelines (jenkins和tekton)
	if config.Pipeline == "jenkins" {
		log.Println("Create jenkins pipeline")
		jauth := jenkins.GetJAuth()
		jks := jenkins.NewJenkins(jauth)
		pipeline, err := util.ReadFile("Jenkinsfile")
		if err != nil {
			log.Fatalln(err)
			return err
		}
		jksconfig := jenkins.CreatePipeline(pipeline)
		err = jks.CreateJob(jksconfig, config.Chart.Name)
		if err != nil {
			log.Println(err)
		}
		// Install helm chart(run jenkins job)
		log.Println("Install helm chart")
		var params map[string]string
		params = make(map[string]string)
		params["action"] = "install"
		params["tag"] = fmt.Sprintf("auto-%s", time.Now().Format("20060102150405"))
		err = jks.BuildWithParams(config.Chart.Name, params)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	} else if config.Pipeline == "tekton" {
		log.Println("Start tekton pipelines")
		res, err := kubernetes.KubeOper("apply", "pipelines/")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(res)
	}

	log.Println("App install success")
	return nil
}

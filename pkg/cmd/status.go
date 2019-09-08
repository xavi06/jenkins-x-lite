package cmd

import (
	"fmt"
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/jenkins"
	"github.com/xavi06/jenkins-x-lite/pkg/tekton"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// Status func
func Status(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Check pipeline
	if config.Pipeline == "jenkins" {
		log.Println("Get job build status")
		jauth := jenkins.GetJAuth()
		jks := jenkins.NewJenkins(jauth)
		isRunning, err := jks.GetBuildStatus(config.Chart.Name)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if isRunning {
			log.Println("Jenkins job is running! Please wait...")
			return nil
		}
		log.Println("Job not running")
	} else if config.Pipeline == "tekton" {
		log.Println("Get tekton taskRun status")
		podStatus, logs, err := tekton.GetTaskRunLog(config)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if podStatus.State == "Running" {
			fmt.Println(logs)
			log.Println("Tekton taskRun is running! Please wait...")
			return nil
		}
		log.Println("Job not running")

	}

	log.Println("Get helm release status")
	printHelmStatus(config.Chart.Name)
	return nil
}

package cmd

import (
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/jenkins"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// JenkinsInit func
func JenkinsInit(c *cli.Context) error {
	log.Println("Init Jenkinsfile")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = jenkins.CreateJenkinsfile(config)
	return err

}

// JenkinsAdd func
func JenkinsAdd(c *cli.Context) error {
	log.Println("Add jenkins job")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
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
		log.Fatalln(err)
		return err
	}
	log.Println("Add success")
	return nil
}

// JenkinsDelete func
func JenkinsDelete(c *cli.Context) error {
	log.Println("Delete jenkins job")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	jauth := jenkins.GetJAuth()
	jks := jenkins.NewJenkins(jauth)
	err = jks.DeleteJob(config.Chart.Name)
	return err

}

// JenkinsUpdate func
func JenkinsUpdate(c *cli.Context) error {
	log.Println("Update jenkins job")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	jauth := jenkins.GetJAuth()
	jks := jenkins.NewJenkins(jauth)
	pipeline, err := util.ReadFile("Jenkinsfile")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	jksconfig := jenkins.CreatePipeline(pipeline)
	err = jks.UpdateJob(jksconfig, config.Chart.Name)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Update success")
	return nil
}

// JenkinsStatus func
func JenkinsStatus(c *cli.Context) error {
	log.Println("Get job build status")
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	jauth := jenkins.GetJAuth()
	jks := jenkins.NewJenkins(jauth)
	isRunning, err := jks.GetBuildStatus(config.Chart.Name)
	if err != nil {
		log.Fatal(err)
		return (err)
	}
	if isRunning {
		log.Println("Jenkins job is running! Please wait...")
		return nil
	}
	log.Println("Job not running")
	return nil

}

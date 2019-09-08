package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/tekton"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// CreatePipelineResource func
func CreatePipelineResource(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create git resource")
	err = tekton.CreateResourceGit(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create image resource")
	_, err = tekton.CreateResourceImage(config)
	if err != nil {
		log.Fatalln(err)
	}
	return nil

}

// CreateTask func
func CreateTask(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create task")
	err = tekton.CreateTask(config, "install", "latest")
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// CreateTaskRun func
func CreateTaskRun(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create taskRun")
	err = tekton.CreateTaskRun(config)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// CreateTekton func
func CreateTekton(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create git resource")
	err = tekton.CreateResourceGit(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create image resource")
	tag, err := tekton.CreateResourceImage(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create task")
	err = tekton.CreateTask(config, "install", tag)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create taskRun")
	err = tekton.CreateTaskRun(config)
	if err != nil {
		log.Fatalln(err)
	}
	return nil

}

// UpdateTekton func
func UpdateTekton(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	os.RemoveAll("./pipelines")
	log.Println("Create git resource")
	err = tekton.CreateResourceGit(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create image resource")
	tag, err := tekton.CreateResourceImage(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create task")
	err = tekton.CreateTask(config, "update", tag)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Create taskRun")
	err = tekton.CreateTaskRun(config)
	if err != nil {
		log.Fatalln(err)
	}
	return nil

}

// GetTektonTaskRunLogs func
func GetTektonTaskRunLogs(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	_, logs, err := tekton.GetTaskRunLog(config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(logs)
	return nil

}

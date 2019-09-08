package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xavi06/jenkins-x-lite/pkg/jenkins"
	"github.com/xavi06/jenkins-x-lite/pkg/kubernetes"
	"github.com/xavi06/jenkins-x-lite/pkg/tekton"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// Update func
func Update(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	if config.Pipeline == "jenkins" {
		log.Println("Rerun jenkins job")
		jauth := jenkins.GetJAuth()
		jks := jenkins.NewJenkins(jauth)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		var params map[string]string
		params = make(map[string]string)
		params["action"] = "update"
		params["tag"] = fmt.Sprintf("auto-%s", time.Now().Format("20060102150405"))
		err = jks.BuildWithParams(config.Chart.Name, params)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		log.Println("Start a build success")
	} else if config.Pipeline == "tekton" {
		log.Println("Rerun tekton pipelines")
		log.Println("Remove old run")
		res, err := kubernetes.KubeOper("delete", "pipelines/")
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
		log.Println("Run a new pipelines")
		res, err = kubernetes.KubeOper("apply", "pipelines/")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(res)
	}
	return nil
}

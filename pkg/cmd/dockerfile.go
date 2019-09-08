package cmd

import (
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/dockerfile"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// DockerfileInit func
func DockerfileInit(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Init Dockerfile")
	dockerf := dockerfile.CreateDockerfileString(config.Docker)
	err = dockerfile.CreateDockerfile(dockerf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Init success")
	return nil
}

// DockerfileUpload func
func DockerfileUpload(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Upload Dockerfile")
	err = dockerfile.UploadDockerfile(config.Chart.Name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Upload success")
	return nil

}

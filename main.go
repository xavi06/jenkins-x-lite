package main

import (
	"log"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/cmd"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "jxl"
	app.Usage = "jenkins x lite, make deploy easy"
	app.Version = "1.2.0"
	app.Commands = cmd.Commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

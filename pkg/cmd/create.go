package cmd

import (
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/config"
	"github.com/urfave/cli"
)

// Create func
func Create(c *cli.Context) error {
	log.Printf("Create %s", configFile)
	var cfg config.Config
	cfg.Name = c.String("name")
	cfg.Apptype = c.String("apptype")
	cfg.Gitrepo = c.String("gitrepo")
	cfg.Gitbranch = c.String("gitbranch")
	cfg.PID = c.Int("pid")
	cfg.Replica = c.Int("replica")
	err := config.CreateToml(configFile, cfg)
	return err
}

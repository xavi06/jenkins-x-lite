package cmd

import (
	"fmt"
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/chartmuseum"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

// UploadChart func
func UploadChart(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Helm package (compress helm chart)")
	chartTgz := fmt.Sprintf("charts/%s-%s.tgz", config.Chart.Name, config.Chart.Version)
	chartDir := fmt.Sprintf("charts/%s", config.Chart.Name)
	err = util.Compress(chartTgz, chartDir)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Upload to chartmuseum")
	chartdata, err := util.ReadFileBytes(chartTgz)
	err = chartmuseum.UploadCharts(chartdata)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Upload success!")
	return nil
}

// DeleteChart func
func DeleteChart(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Delete remote helm chart(chartmuseum)")
	err = chartmuseum.DeleteCharts(config.Chart.Name, config.Chart.Version)
	return err
}

// UpdateChart func
func UpdateChart(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Delete remote helm chart(chartmuseum)")
	err = chartmuseum.DeleteCharts(config.Chart.Name, config.Chart.Version)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Helm package (compress helm chart)")
	chartTgz := fmt.Sprintf("charts/%s-%s.tgz", config.Chart.Name, config.Chart.Version)
	chartDir := fmt.Sprintf("charts/%s", config.Chart.Name)
	err = util.Compress(chartTgz, chartDir)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Upload to chartmuseum")
	chartdata, err := util.ReadFileBytes(chartTgz)
	err = chartmuseum.UploadCharts(chartdata)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Upload success!")
	return nil
}

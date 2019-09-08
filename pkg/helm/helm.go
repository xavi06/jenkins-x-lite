package helm

import (
	"errors"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/util"
	yaml "gopkg.in/yaml.v2"
)

// getTillerHost func Get tiller host
// Default: localhost:44134
func getTillerHost() string {
	tillerAddress := os.Getenv("TILLER_ADDR")
	if tillerAddress == "" {
		tillerAddress = "localhost:44134"
	}
	return tillerAddress

}

// ValuesConfig func
// Export to value.yaml
func ValuesConfig(config util.HelmInfo, fname string) (string, error) {
	yamlconfig, err := yaml.Marshal(&config)
	if err != nil {
		return "", err
	}
	err = util.WriteFileBytes(fname, yamlconfig)
	if err != nil {
		return "", err
	}

	return string(yamlconfig), err
}

// ChartsConfig func
// Export to Charts.yaml
func ChartsConfig(config util.ChartInfo, fname string) (string, error) {
	yamlconfig, err := yaml.Marshal(&config)
	if err != nil {
		return "", err
	}
	err = util.WriteFileBytes(fname, yamlconfig)
	if err != nil {
		return "", err
	}
	return string(yamlconfig), err
}

// InitCharts func
func InitCharts(base, app, url string) error {
	if util.CheckFileExsit("charts") {
		return errors.New("charts exist! Delete it first")
	}
	// download helm chart
	os.Mkdir("charts", os.ModePerm)
	err := util.DownloadFile("charts/charts.tgz", url)
	if err != nil {
		return err
	}
	util.DeCompress("charts/charts.tgz", "charts/")
	os.Remove("charts/charts.tgz")
	os.Rename("charts/"+base, "charts/"+app)
	return nil

}

package chartmuseum

import (
	"fmt"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/util"
)

// GetChartmuseumURL func
func GetChartmuseumURL() string {
	url := os.Getenv("CHARTMUSEUM_URL")
	if url == "" {
		url = "http://localhost:8088"
	}
	return url
}

// UploadCharts func
func UploadCharts(data []byte) error {
	url := GetChartmuseumURL() + "/api/charts"
	_, err := util.DoBytesPost(url, data)
	return err

}

// DeleteCharts func
func DeleteCharts(chart, version string) error {
	var url string
	charturl := GetChartmuseumURL()
	url = fmt.Sprintf("%s/api/charts/%s/%s", charturl, chart, version)
	return util.DoDelete(url)
}

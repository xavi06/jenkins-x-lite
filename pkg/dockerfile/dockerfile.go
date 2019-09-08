package dockerfile

import (
	"errors"
	"os"

	"github.com/xavi06/jenkins-x-lite/pkg/util"
)

// GetFileServer func Get tiller host
// Default: http://localhost:8080
func GetFileServer() string {
	url := os.Getenv("FILESERVER_URL")
	if url == "" {
		url = "http://localhost:8080"
	}
	return url

}

// CreateDockerfileString func
func CreateDockerfileString(docker util.DockerInfo) string {
	var dockerfile string
	dockerfile += "FROM " + docker.Base + "\n"
	dockerfile += "COPY " + docker.Package + " " + docker.Workdir + "\n"
	dockerfile += "WORKDIR " + docker.Workdir + "\n"
	dockerfile += "EXPOSE " + docker.Port + "\n"
	if docker.Entrypoint != "" {
		dockerfile += "ENTRYPOINT " + docker.Entrypoint
	}
	if docker.Cmd != "" {
		dockerfile += "CMD " + docker.Cmd
	}
	return dockerfile
	/*
			return `FROM ` + docker.Base + `
		COPY ` + docker.Package + ` ` + docker.Workdir + `
		WORKDIR ` + docker.Workdir + `
		EXPOSE ` + docker.Port
	*/

}

// CreateDockerfile func
func CreateDockerfile(content string) error {
	if util.CheckFileExsit("Dockerfile") {
		return errors.New("file exist! Delete it first")
	}
	return util.WriteFileString("Dockerfile", content)
}

// UploadDockerfile func
func UploadDockerfile(filename string) error {
	url := GetFileServer()
	url = url + "/upload/dockerfile"
	return util.UploadFile("Dockerfile", filename, url)
}

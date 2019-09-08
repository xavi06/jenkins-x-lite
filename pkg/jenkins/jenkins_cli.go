package jenkins

import (
	"fmt"

	gojks "github.com/xavi06/jenkins_go"
)

// JKS struct
type JKS struct {
	API *gojks.API
}

// NewJenkins func
func NewJenkins(jauth *JAuth) *JKS {
	//jenkins := gojenkins.CreateJenkins(nil, jauth.URL, jauth.User, jauth.Pass)
	jenkins := gojks.NewAPI(jauth.URL, jauth.User, jauth.Pass)
	return &JKS{API: jenkins}
}

// GetJob func
func (jks *JKS) GetJob(job string) (gojks.JobInfo, error) {
	build, err := jks.API.GetJob(job)
	return build, err

}

// CreateJob func
func (jks *JKS) CreateJob(config string, name string) error {
	err := jks.API.CreateJob(name, config)
	return err
}

// BuildWithParams func
func (jks *JKS) BuildWithParams(job string, params map[string]string) error {
	err := jks.API.BuildWithParams(job, params)
	return err
}

// DeleteJob func
func (jks *JKS) DeleteJob(job string) error {
	err := jks.API.DeleteJob(job)
	return err
}

// UpdateJob func
func (jks *JKS) UpdateJob(config, job string) error {
	return jks.API.UpdateJob(job, config)
}

// GetBuildStatus func
func (jks *JKS) GetBuildStatus(job string) (bool, error) {
	Job, err := jks.API.GetJob(job)
	if err != nil {
		return false, err
	}
	LastBuildNo := Job.LastBuild.Number
	Build, err := jks.API.GetJobBuild(job, LastBuildNo)
	if err != nil {
		return false, err
	}
	if Build.Building {
		fmt.Println(jks.API.GetBuildConsoleOutput(job, LastBuildNo))
		return true, nil
	}
	return false, nil
}

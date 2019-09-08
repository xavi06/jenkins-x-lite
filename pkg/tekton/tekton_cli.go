package tekton

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/xavi06/jenkins-x-lite/pkg/kubernetes"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/xavi06/gotekton"
	yaml "gopkg.in/yaml.v2"
)

// CreateResourceGit func
func CreateResourceGit(config *util.TomlConfig) error {
	var params []PipelineResourceParam
	params = []PipelineResourceParam{
		{
			Name:  "revision",
			Value: config.Git.Branch,
		},
		{
			Name:  "url",
			Value: config.Git.Repo,
		},
	}
	yml, err := CreatePipelineResourceString(config.Chart.Name, namespace, "git", params)
	if err != nil {
		return err
	}
	os.Mkdir("./pipelines", os.ModePerm)
	fname := "./pipelines/resource-git.yaml"
	if util.CheckFileExsit(fname) {
		errMsg := fmt.Sprintf("%s exist, please delete it first!", fname)
		return errors.New(errMsg)
	}
	return util.WriteFileString(fname, yml)

}

// CreateResourceImage func
func CreateResourceImage(config *util.TomlConfig) (string, error) {
	var params []PipelineResourceParam
	tag := fmt.Sprintf("auto-%s", time.Now().Format("20060102150405"))
	image := fmt.Sprintf("%s/%s:%s", config.Helm.Image.Repository, config.Helm.Image.Name, tag)
	params = []PipelineResourceParam{
		{
			Name:  "url",
			Value: image,
		},
	}
	yml, err := CreatePipelineResourceString(config.Chart.Name, namespace, "image", params)
	if err != nil {
		return tag, err
	}
	os.Mkdir("./pipelines", os.ModePerm)
	fname := "./pipelines/resource-image.yaml"
	if util.CheckFileExsit(fname) {
		errMsg := fmt.Sprintf("%s exist, please delete it first!", fname)
		return tag, errors.New(errMsg)
	}
	return tag, util.WriteFileString(fname, yml)

}

// CreateTask func
func CreateTask(config *util.TomlConfig, action string, tag string) error {
	// Metadata
	var metaData *MetaData
	var labels map[string]string
	labels = make(map[string]string)
	labels["app"] = config.Chart.Name
	metaData = &MetaData{
		Name:      config.Chart.Name + "-task",
		Namespace: namespace,
		Labels:    labels,
	}
	// Inputs
	inputs := &TaskSpecInputs{
		Resources: []TaskSpecInputsResource{
			{
				Name: "app",
				Type: "git",
			},
		},
		Params: []TaskSpecInputsParam{
			{
				Name:        "pathToDockerFile",
				Description: "The path to the dockerfile to build",
				Default:     "/workspace/app/Dockerfile",
			},
			{
				Name:        "pathToContext",
				Description: "The build context used by Kaniko",
				Default:     "/workspace/app",
			},
		},
	}
	// Outputs
	outputs := &TaskSpecOutputs{
		Resources: []TaskSpecOutputsResource{
			{
				Name: "builtImage",
				Type: "image",
			},
		},
	}
	// Steps
	var steps []TaskSpecStep
	var stepsArr []string
	if action == "install" {
		stepsArr = config.Tekton.InstallSteps
	} else {
		stepsArr = config.Tekton.UpdateSteps
	}
	for _, stepName := range stepsArr {
		steps = append(steps, getStep(stepName, tag, config))
	}
	// Spec
	var spec *TaskSpec
	spec = &TaskSpec{
		Inputs:  *inputs,
		Outputs: *outputs,
		Steps:   steps,
	}
	var task *Task
	task = &Task{
		APIVersion: apiVersion,
		Kind:       "Task",
		MetaData:   *metaData,
		Spec:       *spec,
	}
	yamlconfig, err := yaml.Marshal(&task)
	if err != nil {
		return err
	}
	//fmt.Println(string(yamlconfig))
	os.Mkdir("./pipelines", os.ModePerm)
	fname := "./pipelines/task.yaml"
	if util.CheckFileExsit(fname) {
		errMsg := fmt.Sprintf("%s exist, please delete it first!", fname)
		return errors.New(errMsg)
	}
	return util.WriteFileString(fname, string(yamlconfig))
}

// CreateTaskRun func
func CreateTaskRun(config *util.TomlConfig) error {
	// Metadata
	var metaData *MetaData
	var labels map[string]string
	labels = make(map[string]string)
	labels["app"] = config.Chart.Name
	metaData = &MetaData{
		Name:      config.Chart.Name + "-task-run",
		Namespace: namespace,
		Labels:    labels,
	}
	// Inputs
	inputs := &TaskRunInputs{
		Resources: []TaskRunInputsResource{
			{
				Name:        "app",
				ResourceRef: TaskRunSpecRef{Name: config.Chart.Name + "-git"},
			},
		},
		Params: []PipelineResourceParam{
			{
				Name:  "pathToDockerFile",
				Value: "Dockerfile",
			},
			{
				Name:  "pathToContext",
				Value: "/workspace/app",
			},
		},
	}
	// Outputs
	outputs := &TaskRunOutputs{
		Resources: []TaskRunInputsResource{
			{
				Name:        "builtImage",
				ResourceRef: TaskRunSpecRef{Name: config.Chart.Name + "-image"},
			},
		},
	}
	// Spec
	var spec *TaskRunSpec
	spec = &TaskRunSpec{
		ServiceAccount: serviceAccount,
		TaskRef:        TaskRunSpecRef{Name: config.Chart.Name + "-task"},
		Inputs:         *inputs,
		Outputs:        *outputs,
	}
	var taskRun *TaskRun
	taskRun = &TaskRun{
		APIVersion: apiVersion,
		Kind:       "TaskRun",
		MetaData:   *metaData,
		Spec:       *spec,
	}
	yamlconfig, err := yaml.Marshal(&taskRun)
	if err != nil {
		return err
	}
	//fmt.Println(string(yamlconfig))
	os.Mkdir("./pipelines", os.ModePerm)
	fname := "./pipelines/taskRun.yaml"
	if util.CheckFileExsit(fname) {
		errMsg := fmt.Sprintf("%s exist, please delete it first!", fname)
		return errors.New(errMsg)
	}
	return util.WriteFileString(fname, string(yamlconfig))
}

// GetTaskRunLog func
func GetTaskRunLog(config *util.TomlConfig) (gotekton.PodStatus, string, error) {
	var podStatus gotekton.PodStatus
	kubeconfig := kubernetes.GetKubeconfig()
	client, err := gotekton.GetClientset(kubeconfig)
	if err != nil {
		return podStatus, "", err
	}
	podname := gotekton.GetPodName(client, config.Chart.Name, "tekton-pipelines")
	podStatus, logs, err := gotekton.GetContainerLogsFromPod(client, podname, "tekton-pipelines")
	if err != nil {
		return podStatus, "", err
	}
	return podStatus, logs, err
}

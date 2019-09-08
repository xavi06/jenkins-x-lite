package tekton

// PipelineResource struct
type PipelineResource struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       PipelineResourceSpec
}

// MetaData struct
type MetaData struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

// PipelineResourceSpec struct
type PipelineResourceSpec struct {
	Type   string
	Params []PipelineResourceParam
}

// PipelineResourceParam struct
type PipelineResourceParam struct {
	Name  string
	Value string
}

// Task struct
type Task struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       TaskSpec
}

// TaskSpec struct
type TaskSpec struct {
	Inputs  TaskSpecInputs
	Outputs TaskSpecOutputs
	Steps   []TaskSpecStep
}

// TaskSpecInputs struct
type TaskSpecInputs struct {
	Resources []TaskSpecInputsResource
	Params    []TaskSpecInputsParam
}

// TaskSpecInputsResource struct
type TaskSpecInputsResource struct {
	Name string
	Type string
}

// TaskSpecInputsParam struct
type TaskSpecInputsParam struct {
	Name        string
	Description string
	Default     string
}

// TaskSpecOutputs struct
type TaskSpecOutputs struct {
	Resources []TaskSpecOutputsResource
}

// TaskSpecOutputsResource struct
type TaskSpecOutputsResource struct {
	Name string
	Type string
}

// TaskSpecStep struct
type TaskSpecStep struct {
	Name       string
	Image      string
	Env        []TaskSpecStepsEnv
	WorkingDir string `yaml:"workingDir"`
	Command    []string
	Args       []string
}

// TaskSpecStepsEnv struct
type TaskSpecStepsEnv struct {
	Name  string
	Value string
}

// TaskRun struct
type TaskRun struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       TaskRunSpec
}

// TaskRunSpec struct
type TaskRunSpec struct {
	ServiceAccount string         `yaml:"serviceAccount"`
	TaskRef        TaskRunSpecRef `yaml:"taskRef"`
	Inputs         TaskRunInputs
	Outputs        TaskRunOutputs
}

// TaskRunInputs struct
type TaskRunInputs struct {
	Resources []TaskRunInputsResource
	Params    []PipelineResourceParam
}

// TaskRunSpecRef struct
type TaskRunSpecRef struct {
	Name string
}

// TaskRunInputsResource struct
type TaskRunInputsResource struct {
	Name        string
	ResourceRef TaskRunSpecRef
}

// TaskRunOutputs struct
type TaskRunOutputs struct {
	Resources []TaskRunInputsResource
}

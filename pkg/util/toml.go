package util

import "github.com/BurntSushi/toml"

// TomlConfig struct
type TomlConfig struct {
	Namespace string
	Appbase   string
	Pipeline  string
	Cluster   string
	Helmrepo  string
	Jenkins   JenkinsInfo
	Tekton    TektonInfo
	Chart     ChartInfo
	Git       GitInfo
	Docker    DockerInfo
	Helm      HelmInfo
}

// JenkinsInfo struct
type JenkinsInfo struct {
	Steps        []string
	SyncAddr     []string `toml:"syncAddr" yaml:"syncAddr"`
	GitSecret    string   `toml:"gitSecret" yaml:"gitSecret"`
	DockerSecret string   `toml:"dockerSecret" yaml:"dockerSecret"`
}

// TektonInfo struct
type TektonInfo struct {
	InstallSteps []string `toml:"installSteps" yaml:"installSteps"`
	UpdateSteps  []string `toml:"updateSteps" yaml:"updateSteps"`
}

// ChartInfo struct
type ChartInfo struct {
	APIVersion  string `toml:"apiVersion" yaml:"apiVersion"`
	AppVersion  string `toml:"appVersion" yaml:"appVersion"`
	Name        string
	Version     string
	Description string
}

// HelmInfo struct
type HelmInfo struct {
	Projectid        int
	ReplicaCount     int    `toml:"replicaCount" yaml:"replicaCount"`
	ImagePullSecrets string `toml:"imagePullSecrets" yaml:"imagePullSecrets"`
	Command          []string
	Args             []string
	Image            ImageInfo
	Service          ServiceInfo
	Resources        map[string]ResourceInfo
	Volumes          []map[string]interface{}
	//VolumeMounts     []VolumeMountsInfo `toml:"volumeMounts" yaml:"volumeMounts"`
	VolumeMounts    []map[string]interface{} `toml:"volumeMounts" yaml:"volumeMounts"`
	Env             map[string]EnvInfo
	SecurityContext SecurityContextInfo `toml:"securityContext" yaml:"securityContext"`
	DNSPolicy       string              `toml:"dnsPolicy" yaml:"dnsPolicy"`
	DNSConfig       DNSConfigInfo       `toml:"dnsConfig" yaml:"dnsConfig"`
}

// GitInfo struct
type GitInfo struct {
	Repo   string
	Branch string
}

// DockerInfo struct
type DockerInfo struct {
	Base       string
	Port       string
	Package    string
	Workdir    string
	Entrypoint string
	Cmd        string
}

// ImageInfo struct
type ImageInfo struct {
	Repository string
	Name       string
	Tag        string
	PullPolicy string `toml:"pullPolicy" yaml:"pullPolicy"`
}

// ServiceInfo struct
type ServiceInfo struct {
	Type           string
	Port           int
	TargetPort     int    `toml:"targetport" yaml:"targetport"`
	LoadBalancerIP string `toml:"loadBalancerIP" yaml:"loadBalancerIP"`
	HealthCheck    string `toml:"healthCheck" yaml:"healthCheck"`
}

// ResourceInfo struct
type ResourceInfo struct {
	Memory string
	CPU    string `toml:"cpu" yaml:"cpu"`
}

// VolumeInfo struct
type VolumeInfo struct {
	Name string
	Data map[string]string
}

// VolumeMountsInfo struct
type VolumeMountsInfo struct {
	Name      string
	MountPath string `toml:"mountPath" yaml:"mountPath"`
}

// EnvInfo struct
type EnvInfo struct {
	Name string
	Data map[string]string
}

// SecurityContextInfo struct
type SecurityContextInfo struct {
	RunAsUser int64 `toml:"runAsUser" yaml:"runAsUser"`
	FsGroup   int64 `toml:"fsGroup" yaml:"fsGroup"`
}

// DNSConfigInfo struct
type DNSConfigInfo struct {
	Nameservers []string
}

// ParseConfig func
func ParseConfig(source string) (*TomlConfig, error) {
	var config TomlConfig
	_, err := toml.DecodeFile(source, &config)
	return &config, err

}

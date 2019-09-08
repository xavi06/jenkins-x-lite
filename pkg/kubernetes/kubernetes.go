package kubernetes

import (
	"fmt"
	"os"
	"os/user"
)

// GetKubectl func
/*
func GetKubectl() string {
	os := runtime.GOOS
	switch os {
	case "darwin":
		return "/usr/local/bin/kubectl"
	case "linux":
		return "/usr/bin/kubectl"
	default:
		return "/usr/bin/kubectl"
	}
}
*/

// GetKubeconfig func
func GetKubeconfig() string {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		user, err := user.Current()
		if nil == err {
			return ""
		}
		kubeconfig = fmt.Sprintf("%s/.kube/config", user.HomeDir)
	}
	return kubeconfig
}

package kubernetes

import (
	"fmt"

	"github.com/xavi06/jenkins-x-lite/pkg/util"
)

// KubeOper func
func KubeOper(method string, file string) (string, error) {
	kubectl, err := util.GetExcPath("kubectl")
	if err != nil {
		return "", err
	}
	kubeconfig := GetKubeconfig()
	command := fmt.Sprintf("%s --kubeconfig %s %s -f %s", kubectl, kubeconfig, method, file)
	//fmt.Println(command)
	var args = []string{}
	args = append(args, "-c")
	args = append(args, command)
	outStr, _, err := util.RunCommand("/bin/sh", args...)
	if err != nil {
		return "", err
	}
	return outStr, nil
}

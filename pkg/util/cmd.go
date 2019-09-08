package util

import (
	"bytes"
	"os/exec"
	"path/filepath"
)

// GetExcPath func
func GetExcPath(name string) (string, error) {
	file, err := exec.LookPath(name)
	if err != nil {
		return "", err
	}
	// 获取包含可执行文件名称的路径
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return path, nil
}

// RunCommand func
func RunCommand(name string, args ...string) (outStr, errStr string, err error) {
	cmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return "", "", err
	}
	outStr, errStr = string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr, nil
}

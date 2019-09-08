package util

import yaml "gopkg.in/yaml.v2"

// ConvertToYaml func
func ConvertToYaml(config TomlConfig) ([]byte, error) {
	yamlconfig, err := yaml.Marshal(&config)
	return yamlconfig, err
}

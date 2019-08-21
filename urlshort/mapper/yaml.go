package mapper

import (
	"gopkg.in/yaml.v2"
)

// FromYaml build the mapping from a YAML file
func FromYaml(yamlFile string) (map[string]string, error) {
	return fromFile(yamlFile, yaml.Unmarshal)
}

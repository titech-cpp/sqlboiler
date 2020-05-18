package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// ReadYaml DB設定の書かれたYAMLの読み込み
func ReadYaml(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("Read File Error:%w", err)
	}

	var yamlMap map[string]interface{}
	err = yaml.NewDecoder(file).Decode(&yamlMap)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("Parse Yaml Error:%w", err)
	}

	return yamlMap, nil
}

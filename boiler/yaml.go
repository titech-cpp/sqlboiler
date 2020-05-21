package boiler

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Yaml yamlの構造体
type Yaml struct{}

// ReadYaml DB設定の書かれたYAMLの読み込み
func (*Yaml) ReadYaml(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Read File Error:%w", err)
	}

	err = yaml.NewDecoder(file).Decode(v)
	if err != nil {
		return fmt.Errorf("Parse Yaml Error:%w", err)
	}

	return nil
}

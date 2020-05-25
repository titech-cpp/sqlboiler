package boiler

import (
	"fmt"
	"os"

	"github.com/titech-cpp/sqlboiler/model"
	"gopkg.in/yaml.v2"
)

// Yaml yamlの構造体
type Yaml struct {
	yaml model.Yaml
}

// NewYaml Yamlのコンストラクタ
func NewYaml(path string) (*Yaml,error) {
	boilerYaml := Yaml{
		yaml: model.Yaml{},
	}
	file, err := os.Open(path)
	if err != nil {
		return &Yaml{}, fmt.Errorf("Read File Error:%w", err)
	}

	err = yaml.NewDecoder(file).Decode(&boilerYaml.yaml)
	if err != nil {
		return &Yaml{}, fmt.Errorf("Parse Yaml Error:%w", err)
	}

	return &boilerYaml, nil
}

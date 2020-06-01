package boiler

import (
	"fmt"
	"os"

	"github.com/titech-cpp/sqlboiler/model"
	"gopkg.in/yaml.v2"
)

/*
今の所バリデーションがないのでかなり小さいが、
今後バリデーションがここに入ると他の構造体と同等に肥大化して分離したくなると
予想されるので今のうちから分離している
*/

// Yaml yamlの構造体
type Yaml struct {
	Yaml *model.Yaml
}

// NewYaml Yamlのコンストラクタ
func NewYaml(path string) (*Yaml, error) {
	boilerYaml := Yaml{
		Yaml: &model.Yaml{},
	}
	file, err := os.Open(path)
	if err != nil {
		return &Yaml{}, fmt.Errorf("Read File Error:%w", err)
	}

	err = yaml.NewDecoder(file).Decode(boilerYaml.Yaml)
	if err != nil {
		return &Yaml{}, fmt.Errorf("Parse Yaml Error:%w", err)
	}

	return &boilerYaml, nil
}

func newTestYaml() *Yaml {
	return &Yaml{
		Yaml: &model.Yaml{
			DB: model.DB{
				Type: "mysql",
				Name: "test",
			},
			Tables: map[string][]*model.YamlColumn{
				"test_test": {
					{
						Name: "testa",
						Type:          "int(1)",
						Null:          true,
						AutoIncrement: true,
						Key:           "PRI",
						Default:       "default",
					},
					{
						Name: "testb",
						Type:          "int",
						Null:          false,
						AutoIncrement: false,
						Key:           "",
						Default:       "",
					},
				},
			},
		},
	}
}

package boiler

import (
	"testing"

	"github.com/titech-cpp/sqlboiler/model"
)

func TestReadYaml(t *testing.T) {
	yaml, err := NewYaml("../testdata/test.yaml")
	if err != nil {
		t.Fatalf("ReadYaml Error: %#v", err)
	}
	if err != nil {
		t.Fatalf("Unexpected Yaml Constructor Error: %#v", yaml)
	}

	expect := &model.Yaml{
		DB: model.DB{
			Type: "mysql",
			Name: "test",
		},
		Tables: map[string][]*model.YamlColumn{
			"messages": {
				{
					Name: "user_id",
					Type: "int(11)",
					NoNull: false,
					AutoIncrement: false,
					Key: "",
					Default: "",
				},
			},
		},
	}
	if !yaml.Yaml.Check(expect) {
		t.Fatalf("Invalid Yaml Value %#v, Expected %#v", yaml.Yaml, expect)
	}
}

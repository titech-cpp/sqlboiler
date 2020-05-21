package boiler

import (
	"testing"

	"github.com/titech-cpp/sqlboiler/model"
)

func TestReadYaml(t *testing.T) {
	yaml := new(Yaml)
	var yamls model.Yaml
	err := yaml.ReadYaml("../sample/sqlboiler.yaml", &yamls)
	if err != nil {
		t.Fatalf("ReadYaml Error: %#v", err)
	}

	t.Logf("map: %#v", yamls)
}

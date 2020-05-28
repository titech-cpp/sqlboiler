package boiler

import (
	"testing"
)

func TestReadYaml(t *testing.T) {
	yaml, err := NewYaml("../sample/sqlboiler.yaml")
	if err != nil {
		t.Fatalf("ReadYaml Error: %#v", err)
	}

	t.Logf("map: %#v", yaml.Yaml)
}

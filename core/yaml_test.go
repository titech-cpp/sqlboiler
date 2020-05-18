package core

import (
	"fmt"
	"testing"
)

func checkMap(yamlMap map[string]interface{}, testMap map[string]interface{}) error {
	for key,value := range yamlMap {
		val,ok := testMap[key]
		if !ok {
			return fmt.Errorf("No %s In testMap", key)
		}
		valueMap, ok := value.(map[string]interface{})
		if ok {
			valMap, ok := val.(map[string]interface{})
			if !ok {
				return fmt.Errorf("Unexpected Error: %#v", val)
			}
			err := checkMap(valueMap, valMap)
			return err
		}
		valueString, ok := value.(string)
		if ok {
			valString, ok := val.(string)
			if !ok {
				return fmt.Errorf("Unexpected Error: %#v",val)
			}
			if valueString != valString {
				return fmt.Errorf("Invalid %s", key)
			}
		}
	}
	for key := range testMap {
		_,ok := yamlMap[key]
		if !ok {
			return fmt.Errorf("No %s In yanlMap", key)
		}
	}
	return nil
}

func TestReadYaml(t *testing.T) {
	yamlMap,err := ReadYaml("../testdata/test.yaml")
	if err != nil {
		t.Fatalf("ReadYaml Error: %#v", err)
	}
	testMap := map[string] interface{}{
		"table": map[string]interface{}{
			"messages": map[string]interface{}{
				"user_id": map[string]interface{}{
					"type": "int(11)",
					"null": false,
					"foreign_key": map[string]interface{}{
						"user_id": "id",
					},
				},
			},
		},
	}
	if err = checkMap(yamlMap,testMap); err != nil {
		t.Fatalf("Invalid Map: %#v", err)
	}
	t.Logf("map: %#v", yamlMap)
}
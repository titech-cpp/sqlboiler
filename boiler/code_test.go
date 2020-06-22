package boiler

import (
	"testing"

	"github.com/titech-cpp/sqlboiler/model"
)

func TestCode(t *testing.T) {
	yaml := newTestYaml()
	expect := &model.Code{
		DB: model.DB{
			Type: "mysql",
			Name: "test",
		},
		Tables: []*model.CodeTable{
			{
				Name: &model.NameDetail{
					UpperCamel: "TestTest",
					LowerCamel: "testTest",
					Snake:      "test_test",
				},
				Columns: []*model.CodeColumn{
					{
						Name: &model.NameDetail{
							UpperCamel: "Testa",
							LowerCamel: "testa",
							Snake:      "testa",
						},
						Type:     "nullInt32",
						Null:     true,
						ReadOnly: true,
					},
					{
						Name: &model.NameDetail{
							UpperCamel: "Testb",
							LowerCamel: "testb",
							Snake:      "testb",
						},
						Type:     "int32",
						Null:     false,
						ReadOnly: false,
					},
				},
			},
		},
	}

	code, err := NewCode("test.go", yaml.Yaml)
	if err != nil {
		t.Fatalf("Unexpected Code Constructor Error: %#v", err)
	}

	err = code.Code.Check(expect)
	if err != nil {
		t.Fatalf("Unexpected Invalid Created Code Error: %#v", err)
	}

	expectedType := "bool"
	goType, err := typeParser("boolean", false)
	if err != nil {
		t.Fatalf("Unexpected TypeParser Error: %#v", err)
	}
	if goType != expectedType {
		t.Fatalf("Invalid Parsed Type %s, Expected %s", goType, expectedType)
	}

	_, err = typeParser("test", false)
	if err == nil {
		t.Fatalf("Unexpected No Error: %s", goType)
	}
}

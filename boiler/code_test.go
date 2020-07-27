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
						Type: &model.CodeType{
							Code: &model.CodeSQLTypes{
								NotNull: "nullInt32",
								Null: "int32",
								Upper: "Int32",
							},
							SQL:  "int(1)",
						},
						Null:     true,
						ReadOnly: true,
						Key: &model.CodeKey{
							Primary: true,
							Unique:  false,
							Foreign: []*model.CodeForeign{},
						},
						AutoIncrement: true,
						Default:       "default",
					},
					{
						Name: &model.NameDetail{
							UpperCamel: "Testb",
							LowerCamel: "testb",
							Snake:      "testb",
						},
						Type: &model.CodeType{
							Code: &model.CodeSQLTypes{
								NotNull: "int32",
								Null: "int32",
								Upper: "Int32",
							},
							SQL:  "int",
						},
						Null:     false,
						ReadOnly: false,
						Key: &model.CodeKey{
							Primary: false,
							Unique:  false,
							Foreign: []*model.CodeForeign{},
						},
						AutoIncrement: false,
						Default:       "",
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
	goType, err := typeParser("boolean")
	if err != nil {
		t.Fatalf("Unexpected TypeParser Error: %#v", err)
	}
	if goType.NotNull != expectedType {
		t.Fatalf("Invalid Parsed Type %s, Expected %s", goType, expectedType)
	}

	_, err = typeParser("test")
	if err == nil {
		t.Fatalf("Unexpected No Error: %s", goType)
	}
}

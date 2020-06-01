package boiler

import (
	"testing"

	"github.com/titech-cpp/sqlboiler/model"
)

func TestSchema(t *testing.T) {
	yaml := newTestYaml()
	expect := &model.Schema{
		DB: model.DB{
			Type: "mysql",
			Name: "test",
		},
		Tables: []*model.SchemaTable{
			{
				Name: "test_test",
				Columns: []*model.SchemaColumn{
					{
						Name: "testa",
						Type: "int(1)",
						Null: true,
						Key: "PRI",
						Default: "default",
						Extra: []string{
							"AUTO_INCREMENT",
						},
					},
					{
						Name: "testb",
						Type: "int",
						Null: false,
						Default: "",
						Extra: []string{},
					},
				},
			},
		},
	}

	schema := NewSchema("test.go", yaml.Yaml)
	if !schema.Check(expect) {
		t.Fatalf("Invalid Schema %#v, Expected %#v", schema, expect)
	}
}
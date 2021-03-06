package boiler

import (
	"fmt"

	"github.com/titech-cpp/sqlboiler/boiler/base"
	"github.com/titech-cpp/sqlboiler/model"
)

// Schema スキーマのコンテナ
type Schema struct {
	base.BoilerBase
	*model.Schema
}

// NewSchema Schemaのコンストラクタ
func NewSchema(basePath string, yaml *model.Yaml) *Schema {
	tables := make([]*model.SchemaTable, 0, len(yaml.Tables))
	for key, val := range yaml.Tables {
		columns := make([]*model.SchemaColumn, 0, len(val))
		for _, v := range val {
			extra := make([]string, 0)
			if v.AutoIncrement {
				extra = append(extra, "AUTO_INCREMENT")
			}

			keys := []string{}
			if len(v.Key) != 0 {
				keys = append(keys, v.Key)
			}
			if len(v.ForeignKey) != 0 {
				keys = append(keys, "MUL")
			}

			column := model.SchemaColumn{
				Name:    v.Name,
				Type:    v.Type,
				Null:    !v.NoNull,
				Key:     keys,
				Default: v.Default,
				Extra:   extra,
			}
			columns = append(columns, &column)
		}
		table := model.SchemaTable{
			Name:    key,
			Columns: columns,
		}
		tables = append(tables, &table)
	}
	schemaContainer := model.Schema{
		DB:     yaml.DB,
		Tables: tables,
	}
	boilBase := base.NewBoilBase(basePath)
	schema := Schema{
		BoilerBase: boilBase,
		Schema:     &schemaContainer,
	}
	return &schema
}

// BoilSchema スキーマの生成
func (s *Schema) BoilSchema() error {
	err := s.MakeBaseDir()
	if err != nil {
		return fmt.Errorf("Make Base Directory Error: %w", err)
	}

	fileNames := []string{"dbschema.md"}
	for _, fileName := range fileNames {
		fw, err := s.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		err = s.MakeFile(fw, fileName, s.Schema)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	return nil
}

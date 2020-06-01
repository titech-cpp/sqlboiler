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
		for k, v := range val {
			extra := make([]string, 0)
			if v.AutoIncrement {
				extra = append(extra, "AUTO_INCREMENT")
			}

			column := model.SchemaColumn{
				Name:    k,
				Type:    v.Type,
				Null:    v.Null,
				Key:     v.Key,
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

	fileNames := []string{"dbschemas.go"}
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

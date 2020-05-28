package boiler

import (
	"bytes"
	"fmt"

	"github.com/titech-cpp/sqlboiler/boiler/base"
	"github.com/titech-cpp/sqlboiler/model"
)

// Code コードの構造体
type Code struct {
	base.BoilerBase
	*model.Code
}

// NewCode Codeのコンストラクタ
func NewCode(basePath string, yaml *model.Yaml) (*Code, error) {
	tables := make([]*model.CodeTable, 0, len(yaml.Tables))
	for key, val := range yaml.Tables {
		columns := make([]*model.CodeColumn, 0, len(val))
		for k, v := range val {
			name, err := model.NewNameDetail(k)
			if err != nil {
				return nil, fmt.Errorf("Name Detail Constructor(%s.%s) Error: %w", key, k, err)
			}
			goType, err := typeParser(v.Type)
			if err != nil {
				return nil, fmt.Errorf("Type Parse Error(%s.%s): %w", key, k, err)
			}

			column := &model.CodeColumn{
				Name:     name,
				Type:     goType,
				Null:     v.Null,
				ReadOnly: v.AutoIncrement,
			}
			columns = append(columns, column)
		}
		name, err := model.NewNameDetail(key)
		if err != nil {
			return nil, fmt.Errorf("Name Detail Constructor(%s) Error: %w", key, err)
		}
		table := &model.CodeTable{
			Name:    name,
			Columns: columns,
		}
		tables = append(tables, table)
	}
	codeContainer := &model.Code{
		DB:     yaml.DB,
		Tables: tables,
	}
	boilBase := base.NewBoilBase(basePath)
	code := &Code{
		BoilerBase: boilBase,
		Code:       codeContainer,
	}
	return code, nil
}

// BoilCode コードの生成
func (c *Code) BoilCode() error {
	err := c.MakeFile("tables.go", c.Code)
	if err != nil {
		return fmt.Errorf("Make File Error(%s): %w", "tables.go", err)
	}
	err = c.MakeFile("types.go", c.Code)
	if err != nil {
		return fmt.Errorf("Make File Error(%s): %w", "types.go", err)
	}

	return nil
}

func typeParser(sqlType string) (string, error) {
	typeMap := map[string]string{
		"boolean":   "bool",
		"char":      "string",
		"varchar":   "string",
		"binary":    "string",
		"varbinary": "string",
		"blob":      "blob",
		"text":      "string",
		"integer":   "int32",
		"int":       "int32",
		"bigint":    "int64",
		"mediumint": "int32",
		"smallint":  "int16",
		"tinyint":   "int8",
		"date":      "timeTime",
		"datetime":  "timeTime",
		"timestamp": "timeTime",
		"time":      "timeTime",
		"year":      "timeTime",
	}

	var buf bytes.Buffer
	isStringfied := false
	for _, c := range sqlType {
		if (c < 'a' || 'z' < c) && (c < '0' || '9' < c) && c != '(' && c != ')' {
			return "", fmt.Errorf("%s In Type %s Should Be Lower Case", string(c), sqlType)
		}
		if c == '(' && !isStringfied {
			sqlType = buf.String()
			isStringfied = true
		} else {
			buf.WriteRune(c)
		}
	}
	if !isStringfied {
		sqlType = buf.String()
	}

	goType, ok := typeMap[sqlType]
	if !ok {
		return "", fmt.Errorf("Invalid Type %s", sqlType)
	}

	return goType, nil
}

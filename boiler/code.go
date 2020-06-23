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
	code := new(Code)

	tables := make([]*model.CodeTable, 0, len(yaml.Tables))
	for key, val := range yaml.Tables {
		columns := make([]*model.CodeColumn, 0, len(val))
		for _, v := range val {
			name, err := model.NewNameDetail(v.Name)
			if err != nil {
				return nil, fmt.Errorf("Name Detail Constructor(%s.%s) Error: %w", key, v.Name, err)
			}
			codeType, err := typeParser(v.Type, !v.NoNull)
			if err != nil {
				return nil, fmt.Errorf("Type Parse Error(%s.%s): %w", key, v.Name, err)
			}

			column := &model.CodeColumn{
				Name: name,
				Type: &model.CodeType{
					Code: codeType,
					SQL:  v.Type,
				},
				Null:     !v.NoNull,
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

	code.BoilerBase = boilBase
	code.Code = codeContainer

	return code, nil
}

// BoilCode コードの生成
func (c *Code) BoilCode() error {
	err := c.MakeBaseDir()
	if err != nil {
		return fmt.Errorf("Make Base Directory Error: %w", err)
	}

	fileNames := []string{"tables.go", "types.go", "db.go", "migrate.go"}
	for _, fileName := range fileNames {
		fw, err := c.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		err = c.MakeFile(fw, fileName, c.Code)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	for _, table := range c.Tables {
		fileName := fmt.Sprintf("%s_query.go", table.Name.Snake)
		fw, err := c.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		fileName = "queries.go"
		err = c.MakeFile(fw, fileName, table)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	return nil
}

func typeParser(sqlType string, isNullable bool) (string, error) {
	type sqlTypes struct {
		nonNull string
		null    string
	}
	sqlBool := &sqlTypes{
		nonNull: "bool",
		null:    "nullBool",
	}
	sqlString := &sqlTypes{
		nonNull: "string",
		null:    "nullString",
	}
	sqlInt8 := &sqlTypes{
		nonNull: "int8",
		null:    "nullInt32",
	}
	sqlInt16 := &sqlTypes{
		nonNull: "int16",
		null:    "nullInt32",
	}
	sqlInt32 := &sqlTypes{
		nonNull: "int32",
		null:    "nullInt32",
	}
	sqlInt64 := &sqlTypes{
		nonNull: "int64",
		null:    "nullInt64",
	}
	sqlTime := &sqlTypes{
		nonNull: "timeTime",
		null:    "nullTime",
	}
	typeMap := map[string]*sqlTypes{
		"boolean":   sqlBool,
		"char":      sqlString,
		"varchar":   sqlString,
		"binary":    sqlString,
		"varbinary": sqlString,
		"blob":      sqlString,
		"text":      sqlString,
		"integer":   sqlInt32,
		"int":       sqlInt32,
		"bigint":    sqlInt64,
		"mediumint": sqlInt32,
		"smallint":  sqlInt16,
		"tinyint":   sqlInt8,
		"date":      sqlTime,
		"datetime":  sqlTime,
		"timestamp": sqlTime,
		"time":      sqlTime,
		"year":      sqlTime,
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

	if isNullable {
		return goType.null, nil
	}
	return goType.nonNull, nil
}

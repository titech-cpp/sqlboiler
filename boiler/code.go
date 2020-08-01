package boiler

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/titech-cpp/sqlboiler/boiler/base"
	"github.com/titech-cpp/sqlboiler/model"
)

// Code コードの構造体
type Code struct {
	base.BoilerBase
	base.CopyerBase
	*model.Code
}

// NewCode Codeのコンストラクタ
func NewCode(basePath string, yaml *model.Yaml) (*Code, error) {
	code := new(Code)

	tables := make([]*model.CodeTable, 0, len(yaml.Tables))
	for key, val := range yaml.Tables {
		columns := make([]*model.CodeColumn, 0, len(val))
		joinMap := map[string][]*model.CodeJoinColumn{}
		for _, v := range val {
			name, err := model.NewNameDetail(v.Name)
			if err != nil {
				return nil, fmt.Errorf("Name Detail Constructor(%s.%s) Error: %w", key, v.Name, err)
			}
			codeType, err := typeParser(v.Type)
			if err != nil {
				return nil, fmt.Errorf("Type Parse Error(%s.%s): %w", key, v.Name, err)
			}

			nullCodeType := codeType.Null
			if v.NoNull {
				nullCodeType = codeType.NotNull
			}

			var foreigns []*model.CodeForeign
			for k, yamlForeign := range v.ForeignKey {
				foreign := new(model.CodeForeign)
				foreign.Table = k
				foreign.Column = yamlForeign

				foreigns = append(foreigns, foreign)

				joinTable, err := model.NewNameDetail(k)
				if err != nil {
					return nil, fmt.Errorf("Name Detail Constructor(%s) Error: %w", k, err)
				}

				joinColumn, err := model.NewNameDetail(yamlForeign)
				if err != nil {
					return nil, fmt.Errorf("Name Detail Constructor(%s) Error: %w", k, err)
				}

				joinMap[joinTable.Snake] = append(joinMap[joinTable.Snake], &model.CodeJoinColumn{
					This:   name,
					Target: joinColumn,
				})
			}

			column := &model.CodeColumn{
				Name: name,
				Type: &model.CodeType{
					Code: &model.CodeSQLTypes{
						NotNull: codeType.NotNull,
						Null:    nullCodeType,
						Upper:   codeType.Upper,
					},
					SQL: v.Type,
				},
				Null:     !v.NoNull,
				ReadOnly: v.AutoIncrement,
				Key: &model.CodeKey{
					Primary: v.Key == "PRI",
					Unique:  v.Key == "UNI",
					Foreign: foreigns,
				},
				AutoIncrement: v.AutoIncrement,
				Default:       v.Default,
			}
			columns = append(columns, column)
		}
		name, err := model.NewNameDetail(key)
		if err != nil {
			return nil, fmt.Errorf("Name Detail Constructor(%s) Error: %w", key, err)
		}

		joins := []*model.CodeJoin{}
		for table, column := range joinMap {
			tableName, err := model.NewNameDetail(table)
			if err != nil {
				return nil, fmt.Errorf("Name Detail Constructor(%s) Error: %w", table, err)
			}
			joins = append(joins, &model.CodeJoin{
				Table:  tableName,
				Column: column,
			})
		}
		table := &model.CodeTable{
			Name:    name,
			Columns: columns,
			Joins:   joins,
		}
		tables = append(tables, table)
	}
	codeContainer := &model.Code{
		DB:     yaml.DB,
		Tables: tables,
	}

	boilBase := base.NewBoilBase(basePath)
	copyBase := base.NewCopyBase(basePath + "/query")

	code.BoilerBase = boilBase
	code.CopyerBase = copyBase
	code.Code = codeContainer
	code.makeJoinedTables()

	return code, nil
}

// BoilCode コードの生成
func (c *Code) BoilCode() error {
	err := c.BoilerBase.MakeBaseDir()
	if err != nil {
		return fmt.Errorf("Make Base Directory Error(Boiler): %w", err)
	}

	err = c.CopyerBase.MakeBaseDir()
	if err != nil {
		return fmt.Errorf("Make Base Directory Error(Copy): %w", err)
	}

	fileNames := []string{"types.go", "db.go", "migrate.go"}
	for _, fileName := range fileNames {
		fw, err := c.BoilerBase.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		err = c.BoilerBase.MakeFile(fw, fileName, c.Code)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	fileNames = []string{"migrate", "where", "join"}
	for _, fileName := range fileNames {
		fw, err := c.CopyerBase.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		err = c.CopyerBase.MakeFile(fw, fileName)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	for _, table := range c.JoinedTables {
		strTables := make([]string, 0, len(table.Tables))
		for _, v := range table.Tables {
			strTables = append(strTables, v.Name.Snake)
		}

		fileName := fmt.Sprintf("join_%s_query.go", strings.Join(strTables, "_"))
		fw, err := c.BoilerBase.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		fileName = "joined_queries.go"
		err = c.BoilerBase.MakeFile(fw, fileName, table)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	for _, table := range c.Tables {
		fileName := fmt.Sprintf("%s_query.go", table.Name.Snake)
		fw, err := c.BoilerBase.MakeFileWriter(fileName)
		if err != nil {
			return fmt.Errorf("Make File Writer Error(%s): %w", fileName, err)
		}
		fileName = "queries.go"
		err = c.BoilerBase.MakeFile(fw, fileName, table)
		if err != nil {
			return fmt.Errorf("Make File Error(%s): %w", fileName, err)
		}
	}

	return nil
}

func (c *Code) makeJoinedTables() error {
	joinedTables := []*model.CodeJoinedTable{}

	tableMap := map[string]*model.CodeTable{}
	for _, table := range c.Tables {
		tableMap[table.Name.Snake] = table
	}

	for _, table := range c.Tables {
		for _, join := range table.Joins {
			joinedTable, ok := tableMap[join.Table.Snake]
			if !ok {
				return fmt.Errorf("invalid join table: %s", join.Table)
			}

			joinedJoins := []*model.CodeJoin{}
			for _, v := range append(table.Joins, joinedTable.Joins...) {
				if v != join {
					joinedJoins = append(joinedJoins, v)
				}
			}

			joinedTables = append(joinedTables, &model.CodeJoinedTable{
				Tables: []*model.CodeTable{
					table,
					joinedTable,
				},
				Joins: joinedJoins,
			})
		}
	}

	c.JoinedTables = append(c.JoinedTables, joinedTables...)

	return nil
}

func (c *Code) makeDoubleJoinedTables(tables []*model.CodeJoinedTable, tableMap map[string]*model.CodeTable) error {
	joinedTables := []*model.CodeJoinedTable{}

	for _, table := range tables {
		for _, join := range table.Joins {
			joinedJoins := []*model.CodeJoin{}
			for _, v := range table.Joins {
				if v != join {
					joinedJoins = append(joinedJoins, v)
				}
			}

			joinedTable, ok := tableMap[join.Table.Snake]
			if !ok {
				return fmt.Errorf("invalid join table: %s", join.Table)
			}

			joinedTables = append(joinedTables, &model.CodeJoinedTable{
				Tables: append(table.Tables, joinedTable),
				Joins:  joinedJoins,
			})
		}
	}

	c.JoinedTables = append(c.JoinedTables, joinedTables...)

	if len(joinedTables) != 0 {
		err := c.makeDoubleJoinedTables(joinedTables, tableMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func typeParser(sqlType string) (*model.CodeSQLTypes, error) {
	sqlBool := &model.CodeSQLTypes{
		NotNull: "bool",
		Null:    "nullBool",
		Upper:   "Bool",
	}
	sqlString := &model.CodeSQLTypes{
		NotNull: "string",
		Null:    "nullString",
		Upper:   "String",
	}
	sqlInt8 := &model.CodeSQLTypes{
		NotNull: "int8",
		Null:    "nullInt32",
		Upper:   "Int32",
	}
	sqlInt16 := &model.CodeSQLTypes{
		NotNull: "int16",
		Null:    "nullInt32",
		Upper:   "Int32",
	}
	sqlInt32 := &model.CodeSQLTypes{
		NotNull: "int32",
		Null:    "nullInt32",
		Upper:   "Int32",
	}
	sqlInt64 := &model.CodeSQLTypes{
		NotNull: "int64",
		Null:    "nullInt64",
		Upper:   "Int64",
	}
	sqlTime := &model.CodeSQLTypes{
		NotNull: "timeTime",
		Null:    "nullTime",
		Upper:   "Time",
	}
	typeMap := map[string]*model.CodeSQLTypes{
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
			return nil, fmt.Errorf("%s In Type %s Should Be Lower Case", string(c), sqlType)
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
		return nil, fmt.Errorf("Invalid Type %s", sqlType)
	}

	return goType, nil
}

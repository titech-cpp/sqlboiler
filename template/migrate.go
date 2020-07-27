package query

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
)

// Table テーブルの構造体
type Table []*Column

// Column カラムの構造体
type Column struct {
	Field   string
	Type    string
	Null    bool
	Key     []string
	Default string
	Extra   []string
}

// Key キーの構造体
type Key struct {
	Name   string
	Column *Column
}

// ShowDatabaseInfo データベースの情報の取得
func ShowDatabaseInfo(ctx context.Context, tx *sql.Tx) (tableMap map[string]Table, err error) {
	tableMap = map[string]Table{}
	tableNames := []string{}
	rows, err := tx.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		return nil, fmt.Errorf("Get Tables Error: %w", err)
	}
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, fmt.Errorf("Scan Tables Error: %w", err)
		}
		tableNames = append(tableNames, tableName)
	}

	for _, tableName := range tableNames {
		rows, err := tx.QueryContext(ctx, "SHOW COLUMNS FROM "+tableName)
		if err != nil {
			return nil, fmt.Errorf("Get Columns Error: %w", err)
		}
		columns := Table{}
		rep := regexp.MustCompile("%s*,%s*")
		for rows.Next() {
			var column Column
			var null string
			var key sql.NullString
			var defaultVal sql.NullString
			var extra sql.NullString
			err := rows.Scan(&column.Field, &column.Type, &null, &key, &defaultVal, &extra)
			if err != nil {
				return nil, fmt.Errorf("Scan Columns Error: %w", err)
			}

			if null == "YES" {
				column.Null = true
			} else {
				column.Null = false
			}

			var keys []string
			if key.Valid {
				keys = rep.Split(key.String, -1)
			}
			column.Key = []string{}
			for _, v := range keys {
				if v != "MUL" {
					column.Key = append(column.Key, v)
				}
			}

			if defaultVal.Valid {
				column.Default = defaultVal.String
			}

			column.Extra = []string{}
			if extra.Valid {
				column.Extra = rep.Split(extra.String, -1)
			}

			columns = append(columns, &column)
		}
		tableMap[tableName] = columns
	}

	return
}

// GetDiff テーブル間の差分をとる
func (nowTable Table) GetDiff(newTable Table) (addColumns []*Column, alterColumns []*Column, changeNullColumns []*Column, delColumns []*Column, addKeys []*Key, delKeys []*Key) {
	for _, column := range nowTable {
		for i, v := range newTable {
			if column.Field == v.Field {
				if column.Type != v.Type {
					fmt.Println("type", *v)
					alterColumns = append(alterColumns, v)
				}
				if column.Null != v.Null {
					fmt.Println("null", column, *v)
					changeNullColumns = append(changeNullColumns, v)
				}
				for _, key := range column.Key {
					for i, val := range v.Key {
						if key == val {
							break
						}
						if i == len(v.Key)-1 {
							fmt.Println("add_key", *v)
							addKeys = append(addKeys, &Key{
								Name:   key,
								Column: v,
							})
						}
					}
				}
				for _, key := range v.Key {
					for i, val := range column.Key {
						if key == val {
							break
						}
						if i == len(v.Key)-1 {
							fmt.Println("del_key", *v)
							delKeys = append(delKeys, &Key{
								Name:   key,
								Column: v,
							})
						}
					}
				}
				for _, extra := range column.Extra {
					for i, val := range v.Extra {
						if extra == val {
							break
						}
						if i == len(v.Extra) - 1 {
							fmt.Println("del_extra", *v)
							alterColumns = append(alterColumns, v)
						}
					}
				}
				for _, extra := range v.Extra {
					for i, val := range column.Extra {
						if extra == val {
							break
						}
						if i == len(v.Extra) - 1 {
							fmt.Println("add_extra", *v)
							alterColumns = append(alterColumns, v)
						}
					}
				}
				break
			}
			if i == len(newTable)-1 {
				fmt.Println("del_column", *v)
				delColumns = append(delColumns, column)
			}
		}
	}
	for _, column := range newTable {
		for i, v := range nowTable {
			if column.Field == v.Field {
				break
			}
			if i == len(nowTable)-1 {
				fmt.Println("add_column", *v)
				addColumns = append(addColumns, column)
			}
		}
	}

	return
}

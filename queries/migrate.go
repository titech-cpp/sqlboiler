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
	Null    string
	Key     []string
	Default string
	Extra   []string
}

// Key キーの構造体
type Key struct {
	Name   string
	Column string
	Type   string
}

// ShowDatabaseInfo データベースの情報の取得
func ShowDatabaseInfo(ctx context.Context, tx *sql.Tx) (tableMap map[string]Table, err error) {
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
		rows, err := tx.QueryContext(ctx, "SHOW COLUMNS FROM " + tableName)
		if err != nil {
			return nil, fmt.Errorf("Get Columns Error: %w", err)
		}
		columns := Table{}
		rep := regexp.MustCompile("%s*,%s*")
		for rows.Next() {
			var column Column
			var key string
			var extra string
			err := rows.Scan(&column.Field, &column.Type, &column.Null, &key, &column.Default, &extra)
			if err != nil {
				return nil, fmt.Errorf("Scan Columns Error: %w", err)
			}

			keys := rep.Split(key, -1)
			column.Key = []string{}
			for _, v := range keys {
				if v != "MUL" {
					column.Key = append(column.Key, v)
				}
			}

			column.Extra = rep.Split(extra, -1)

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
					alterColumns = append(alterColumns, v)
				}
				if column.Null != v.Null {
					changeNullColumns = append(changeNullColumns, v)
				}
				for _, key := range column.Key {
					for i, val := range v.Key {
						if key == val {
							break
						}
						if i == len(v.Key)-1 {
							addKeys = append(addKeys, &Key{
								Name:   key,
								Column: column.Field,
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
							delKeys = append(delKeys, &Key{
								Name:   key,
								Column: column.Field,
							})
						}
					}
				}
				break
			}
			if i == len(newTable)-1 {
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
				addColumns = append(addColumns, column)
			}
		}
	}

	return
}

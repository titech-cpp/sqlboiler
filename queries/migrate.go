package query

import (
	"context"
	"database/sql"
	"fmt"
)

// Column カラムの構造体
type Column struct {
	Field string
	Type string
	Null string
	Key string
	Extra string
}

func (c *Column) equal(clm *Column) bool {
	return c.Field == clm.Field && c.Type == clm.Type && c.Null == clm.Null && c.Key == clm.Key && c.Extra == clm.Extra
}

// ShowDatabaseInfo データベースの情報の取得
func ShowDatabaseInfo(ctx context.Context, tx *sql.Tx) (tableMap map[string][]*Column, err error) {
	tableNames := []string{}
    rows,err := tx.QueryContext(ctx, "SHOW TABLES")
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

    for _,tableName := range tableNames {
		rows,err := tx.QueryContext(ctx, "SHOW COLUMNS FROM ?", tableName)
		if err != nil {
			return nil, fmt.Errorf("Get Columns Error: %w", err)
		}
		columns := []*Column{}
		for rows.Next() {
			var column Column
			err := rows.Scan(&column.Field, &column.Type, &column.Null, &column.Key, &column.Extra)
			if err != nil {
				return nil, fmt.Errorf("Scan Columns Error: %w", err)
			}
			columns = append(columns, &column)
		}
		tableMap[tableName] = columns
	}

	return
}
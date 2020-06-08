package models

import (
    "fmt"
    "database/sql"

    "github.com/titech-cpp/sqlboiler/queries"
)

// {{.Name.UpperCamel}}Query {{.Name.UpperCamel}}のクエリの構造体
type {{.Name.UpperCamel}}Query struct {
    db *sql.DB
    *query.Where{{range .Columns}}
    {{.Name.UpperCamel}} interface{}{{end}}
}

func (q *{{.Name.UpperCamel}}Query) createWhereMap() (whereMap map[string]interface{}) { {{range .Columns}}
    if q.{{.Name.UpperCamel}} != nil {
        whereMap["{{.Name.Snake}}"] = q.{{.Name.UpperCamel}}
    }{{end}}
    return
}

func (q *{{.Name.UpperCamel}}Query) Find() (*{{.Name.UpperCamel}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.Where.Where(whereMap)

    args := []interface{}{}
    args = append(args, whereArgs)
    query := fmt.Sprintf("SELECT * FROM {{.Name.Snake}} LIMIT 1 %s", whereQuery)
    rows, err := q.db.Query(query, args)
    if err != nil {
        return &{{.Name.UpperCamel}}Table{}, err
    }

    table := new({{.Name.UpperCamel}}Table)
    dests := []interface{}{}{{range .Columns}}
    dests = append(dests, &(table.{{.Name.UpperCamel}})){{end}}
    rows.Next()
    err = rows.Scan(dests...)
    if err != nil {
        return &{{.Name.UpperCamel}}Table{}, err
    }

    return table, nil
}
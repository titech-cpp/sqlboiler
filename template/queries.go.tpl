package models

import (
    "fmt"
    "database/sql"

    "github.com/titech-cpp/sqlboiler/queries"
)

type {{.Name.LowerCamel}} int
{{$lowerTable := .Name.LowerCamel}}{{$upperTable := .Name.UpperCamel}}
const ({{range $i, $v := .Columns}}
    {{printf $upperTable}}{{.Name.UpperCamel}}{{if eq $i 0}} {{printf $lowerTable}} = iota{{end}}{{end}}
)

// {{.Name.UpperCamel}}Query {{.Name.UpperCamel}}のクエリの構造体
type {{.Name.UpperCamel}}Query struct {
    db *sql.DB
    whereStruct *query.Where{{range .Columns}}
    {{.Name.UpperCamel}} interface{}{{end}}
}

func (q *{{.Name.UpperCamel}}Query) createWhereMap() (whereMap map[string]interface{}) { {{range .Columns}}
    if q.{{.Name.UpperCamel}} != nil {
        whereMap["{{.Name.Snake}}"] = q.{{.Name.UpperCamel}}
    }{{end}}
    return
}

func (q *{{.Name.UpperCamel}}Query) Where(columnType {{printf $lowerTable}}, value interface{}) *{{.Name.UpperCamel}}Query {
    switch columnType { {{range .Columns}}
        case {{printf $upperTable}}{{.Name.UpperCamel}}:
            q.{{.Name.UpperCamel}} = value{{end}}
    }

    return q
}

func (q *{{.Name.UpperCamel}}Query) Find() (*{{.Name.UpperCamel}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

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
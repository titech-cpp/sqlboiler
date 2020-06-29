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

{{template "select" .}}

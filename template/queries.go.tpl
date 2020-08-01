package models

import (
    "fmt"
    "strings"
    "database/sql"

    "sample/models/query"
)

type {{.Name.LowerCamel}} int
{{$lowerTable := .Name.LowerCamel}}{{$upperTable := .Name.UpperCamel}}
const ({{range $i, $v := .Columns}}
    {{printf $upperTable}}{{.Name.UpperCamel}}{{if eq $i 0}} {{printf $lowerTable}} = iota{{end}}{{end}}
)

// {{.Name.UpperCamel}}Query {{.Name.UpperCamel}}のクエリの構造体
type {{.Name.UpperCamel}}Query struct {
    db *sql.DB
    whereStruct *query.Where
    table {{.Name.UpperCamel}}
}

{{template "where" .}}

{{template "select" .}}

{{template "insert" .}}

{{template "update" .}}

{{template "delete" .}}

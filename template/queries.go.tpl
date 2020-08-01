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

// {{.Name.UpperCamel}} {{.Name.UpperCamel}}のテーブルの構造体（全ての要素がポインタ）
type {{.Name.UpperCamel}} struct {
{{range .Columns}}    {{.Name.UpperCamel}} *{{.Type.Code.NotNull}}
{{end}}}

type nullable{{.Name.UpperCamel}}Table struct {
{{range .Columns}}    {{.Name.UpperCamel}} {{.Type.Code.Null}}
{{end}}}

// {{.Name.UpperCamel}}Table {{.Name.UpperCamel}}のテーブルの構造体
type {{.Name.UpperCamel}}Table struct {
{{range .Columns}}    {{.Name.UpperCamel}} {{.Type.Code.NotNull}}
{{end}}}

// {{.Name.UpperCamel}}Query {{.Name.UpperCamel}}のクエリの構造体
type {{.Name.UpperCamel}}Query struct {
    db *sql.DB
    whereStruct *query.Where
    limit int
    offset int
    table {{.Name.UpperCamel}}
}

{{template "where" .}}

{{template "select" .}}

{{template "insert" .}}

{{template "update" .}}

{{template "delete" .}}

{{template "join" .}}

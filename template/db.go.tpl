package models

import (
    "database/sql"
    "errors"
)

var (
    RECORD_NOT_FOUND = errors.New("record not found")
)

type DB struct {
    db *sql.DB
}

func NewDB(sqlDB *sql.DB) *DB {
    db := new(DB)
    db.db = sqlDB
    return db
}{{range .Tables}}

func (d *DB) {{.Name.UpperCamel}}() *{{.Name.UpperCamel}}Query {
    table := new({{.Name.UpperCamel}}Query)
    table.db = d.db

    return table
}{{end}}

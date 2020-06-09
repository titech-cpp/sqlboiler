package models

import (
    "database/sql"
)

type DB struct {
    db *sql.DB
}

func NewDB(sqlDB *sql.DB) *DB {
    db := new(DB)
    db.db = sqlDB
    return db
}{{range .Tables}}

func (*DB) {{.Name.UpperCamel}}() *{{.Name.UpperCamel}}Query {
    return new({{.Name.UpperCamel}}Query)
}{{end}}

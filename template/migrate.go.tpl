package models

import (
    "fmt"
    "context"
)

func (d *DB) Migrate() error {
    var cmd string
    var err error

    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    db := d.db
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("Start Transaction Error: %w", err)
    }

    {{range .Tables}}{{ $length := len .Columns }}
    cmd = `CREATE TABLE IF NOT EXISTS {{.Name.Snake}}({{range $i, $v := .Columns}}
        {{$v.Name.Snake}} {{$v.Type.SQL}}{{if not $v.Null}} NOT NULL{{end}}{{if .Key.Primary}} PRIMARY KEY{{end}}{{if .Key.Unique}} UNIQUE{{end}}{{if .AutoIncrement}} AUTO_INCREMENT{{end}}{{ $default := .Default }}{{if ne .Default ""}} DEFAULT {{$default}}{{end}}{{if not (isLast $i $length)}},{{end}}{{end}}
    )`
    _, err = tx.ExecContext(ctx, cmd)
    if err != nil {
        return fmt.Errorf("Create Table Error(Table {{.Name.Snake}}): %w", err)
    }
    {{end}}
    {{range .Tables}}{{ $table := .Name.Snake}}{{range .Columns}}{{ $column := .Name.Snake}}{{range .Key.Foreign}}
    cmd = `ALTER TABLE {{printf $table}} ADD FOREIGN KEY ({{printf $column}}) REFERENCES {{.Table}}({{.Column}})`
    _, err = tx.ExecContext(ctx, cmd)
    if err != nil {
        return fmt.Errorf("Add Foreign Key Error({{printf $table}}.{{printf $column}}): %w", err)
    }
    {{end}}{{end}}{{end}}
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("Trasction Commit Error: %w", err)
    }

    return nil
}
package models

import "fmt"

func (d *DB) Migrate() error {
    var cmd string
    var err error
    db := d.db
    {{range .Tables}}{{ $length := len .Columns }}cmd = `CREATE TABLE IF NOT EXISTS {{.Name.Snake}}({{range $i, $v := .Columns}}
        {{$v.Name.Snake}} {{$v.Type.SQL}}{{if not $v.Null}} NOT NULL{{end}}{{if not (isLast $i $length)}},{{end}}{{end}}
    )`
    _, err = db.Exec(cmd)
    if err != nil {
        return fmt.Errorf("Create Table Error(Table {{.Name.Snake}}): %w", err)
    }
    {{end}}
    return nil
}
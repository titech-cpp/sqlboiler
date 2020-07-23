package models

import (
    "fmt"
    "context"

    "sample/models/query"
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

    nowTableMap, err := query.ShowDatabaseInfo(ctx, tx)
    if err != nil {
        return fmt.Errorf("Get Database Infomation Error: %w", err)
    }

    newTableMap := map[string]query.Table{
{{range .Tables}}        "{{.Name.Snake}}": {
{{range .Columns}}            &query.Column{
                Field: "{{.Name.Snake}}",
                Type: "{{.Type.SQL}}",
                Null: {{.Null}},
                Key: []string{
{{if .Key.Primary}}                    "PRI",
{{end}}{{if .Key.Unique}}                    "UNI",
{{end}}                },
                Default: "{{.Default}}",
                Extra: []string{
{{if .AutoIncrement}}                    "auto_increment",
{{end}}                },
            },
{{end}}        },
{{end}}    }

    cmds := []string{}
    for k,v := range newTableMap {
        t, ok := nowTableMap[k]
        if !ok {
            cmd := "CREATE TABLE IF NOT EXISTS " + k + " ("
            for i, val := range v {
                null := ""
                if !val.Null {
                    null = " NOT NULL"
                }
                key := ""
                for _,value := range val.Key {
                    switch value {
                    case "PRI":
                        key += " PRIMARY KEY"
                    case "UNI":
                        key += " UNIQUE"
                    }
                }
                autoIncrement := ""
                for _, value := range val.Extra {
                    switch value {
                    case "auto_increment":
                        autoIncrement = "AUTO_INCREMENT"
                    }
                }
                defaultVal := ""
                if len(val.Default) != 0 {
                    defaultVal = " DEFAULT " + val.Default
                }
                cmd += fmt.Sprintf("%s %s%s%s%s%s", val.Field, val.Type, null, key, autoIncrement, defaultVal)
                if i != len(v)-1 {
                    cmd += ","
                } else {
                    cmd += ")"
                }
            }
            cmds = append(cmds, cmd)
        } else {
            addColumns, alterColumns, changeNullColumns, delColumns, addKeys, delKeys := t.GetDiff(v)

            for _,val := range addColumns {
                null := ""
                if !val.Null {
                    null = " NOT NULL"
                }
                cmd := fmt.Sprintf("ALTER TABLE %s ADD %s %s%s", k, val.Field, val.Type, null)
                cmds = append(cmds, cmd)
            }

            for _, val := range delColumns {
                cmd := fmt.Sprintf("ALTER TABLE %s DROP %s", k, val.Field)
                cmds = append(cmds, cmd)
            }

            changeKeyColumns := []*query.Column{}
            for _, val := range append(addKeys, delKeys...) {
                changeKeyColumns = append(changeKeyColumns, val.Column)
            }

            for _, val := range append(append(alterColumns, changeNullColumns...), changeKeyColumns...) {
                null := ""
                if !val.Null {
                    null = " NOT NULL"
                }
                key := ""
                for _,value := range val.Key {
                    switch value {
                    case "PRI":
                        key += " PRIMARY KEY"
                    case "UNI":
                        key += " UNIQUE"
                    }
                }
                autoIncrement := ""
                for _, value := range val.Extra {
                    switch value {
                    case "auto_increment":
                        autoIncrement = " AUTO_INCREMENT"
                    }
                }
                defaultVal := ""
                if len(val.Default) != 0 {
                    defaultVal = " DEFAULT " + val.Default
                }
                columnDefinition := fmt.Sprintf("%s %s%s%s%s%s", val.Field, val.Type, null, key, autoIncrement, defaultVal)

                cmd := fmt.Sprintf("ALTER TABLE %s MODIFY %s", k, columnDefinition)
                cmds = append(cmds, cmd)
            }
        }
    }

    for _,cmd := range cmds {
        _, err = tx.ExecContext(ctx, cmd)
        if err != nil {
            return fmt.Errorf("Create Table Error(%s): %w", cmd, err)
        }
    }

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
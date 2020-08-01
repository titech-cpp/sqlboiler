{{define "joinedSelect"}}{{$l := len .Tables}}
func (q *Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query) Find() (*Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }

    queryTable := "{{range $i,$v := .Tables}}{{if ne $i 0}} " + (*(q.joins[{{printf "%d" $i}}-1]))[0] + " {{end}}{{$v.Name.Snake}}{{if ne $i 0}} " + (*(q.joins[{{printf "%d" $i}}-1]))[1] + " {{end}}{{end}}"

    query := fmt.Sprintf("SELECT * FROM %s %s LIMIT 1", queryTable, whereQuery)

    fmt.Println(query)
    rows, err := q.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("Execute Query Error: %w", err)
    }
    defer rows.Close()

    nullableTable := new(nullableJoin{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table)
    dests := []interface{}{}{{range .Tables}}{{$upperTable := .Name.UpperCamel}}{{range .Columns}}
    dests = append(dests, &(nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}})){{end}}{{end}}
    if rows.Next() {
        err = rows.Scan(dests...)
        if err != nil {
            return nil, fmt.Errorf("Scaning Rows Error: %w", err)
        }
    } else {
        return nil, RECORD_NOT_FOUND
    }

    table := new(Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table){{range .Tables}}{{$upperTable := .Name.UpperCamel}}{{range .Columns}}{{if .Null}}
    if nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}}.Valid {
        table.{{printf $upperTable}}.{{.Name.UpperCamel}} = nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}}.{{.Type.Code.Upper}}
    }{{else}}
    table.{{printf $upperTable}}.{{.Name.UpperCamel}} = nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}}{{end}}{{end}}{{end}}

    return table, nil
}

func (q *Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query) Select() ([]*Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }

    queryTable := "{{range $i,$v := .Tables}}{{if ne $i 0}} " + (*(q.joins[{{printf "%d" $i}}-1]))[0] + " {{end}}{{$v.Name.Snake}}{{if ne $i 0}} " + (*(q.joins[{{printf "%d" $i}}-1]))[1] + " {{end}}{{end}}"

    query := fmt.Sprintf("SELECT * FROM %s %s", queryTable, whereQuery)

    fmt.Println(query)
    rows, err := q.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("Execute Query Error: %w", err)
    }
    defer rows.Close()

    tables := []*Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table{}
    for rows.Next() {
        nullableTable := new(nullableJoin{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table)
        dests := []interface{}{}{{range .Tables}}{{$upperTable := .Name.UpperCamel}}{{range .Columns}}
        dests = append(dests, &(nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}})){{end}}{{end}}
        err = rows.Scan(dests...)
        if err != nil {
            return nil, fmt.Errorf("Scaning Rows Error: %w", err)
        }

        table := new(Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table){{range .Tables}}{{$upperTable := .Name.UpperCamel}}{{range .Columns}}{{if .Null}}
        if nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}}.Valid {
            table.{{printf $upperTable}}.{{.Name.UpperCamel}} = nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}}.{{.Type.Code.Upper}}
        }{{else}}
        table.{{printf $upperTable}}.{{.Name.UpperCamel}} = nullableTable.{{printf $upperTable}}.{{.Name.UpperCamel}}{{end}}{{end}}{{end}}

        tables = append(tables, table)
    }

    return tables, nil
}
{{end}}
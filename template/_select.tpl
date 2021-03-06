{{define "select"}}
func (q *{{.Name.UpperCamel}}Query) Limit(limit int) *{{.Name.UpperCamel}}Query {
    q.limit = limit

    return q
}

func (q *{{.Name.UpperCamel}}Query) Offset(offset int) *{{.Name.UpperCamel}}Query {
    q.offset = offset

    return q
}

func (q *{{.Name.UpperCamel}}Query) Find() (*{{.Name.UpperCamel}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }
    query := fmt.Sprintf("SELECT * FROM {{.Name.Snake}} %s LIMIT 1", whereQuery)

    rows, err := q.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("Execute Query Error: %w", err)
    }
    defer rows.Close()

    nullableTable := new(nullable{{.Name.UpperCamel}}Table)
    dests := []interface{}{}{{range .Columns}}
    dests = append(dests, &(nullableTable.{{.Name.UpperCamel}})){{end}}
    if rows.Next() {
        err = rows.Scan(dests...)
        if err != nil {
            return nil, fmt.Errorf("Scaning Rows Error: %w", err)
        }
    } else {
        return nil, RECORD_NOT_FOUND
    }

    table := new({{.Name.UpperCamel}}Table){{range .Columns}}{{if .Null}}
    if nullableTable.{{.Name.UpperCamel}}.Valid {
        table.{{.Name.UpperCamel}} = nullableTable.{{.Name.UpperCamel}}.{{.Type.Code.Upper}}
    }{{else}}
    table.{{.Name.UpperCamel}} = nullableTable.{{.Name.UpperCamel}}{{end}}{{end}}

    return table, nil
}

func (q *{{.Name.UpperCamel}}Query) Select() ([]*{{.Name.UpperCamel}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }

    query := fmt.Sprintf("SELECT * FROM {{.Name.Snake}} %s", whereQuery)

    if q.limit != 0 {
        query += " LIMIT ?"
        args = append(args, q.limit)
    }

    if q.offset != 0 {
        query += " OFFSET ?"
        args = append(args, q.offset)
    }

    rows, err := q.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("Execute Query Error: %w", err)
    }
    defer rows.Close()

    tables := []*{{.Name.UpperCamel}}Table{}
    for rows.Next() {
        nullableTable := new(nullable{{.Name.UpperCamel}}Table)
        dests := []interface{}{}{{range .Columns}}
        dests = append(dests, &(nullableTable.{{.Name.UpperCamel}})){{end}}
        err = rows.Scan(dests...)
        if err != nil {
            return nil, fmt.Errorf("Scaning Rows Error: %w", err)
        }

        table := new({{.Name.UpperCamel}}Table){{range .Columns}}{{if .Null}}
        if nullableTable.{{.Name.UpperCamel}}.Valid {
            table.{{.Name.UpperCamel}} = nullableTable.{{.Name.UpperCamel}}.{{.Type.Code.Upper}}
        }{{else}}
        table.{{.Name.UpperCamel}} = nullableTable.{{.Name.UpperCamel}}{{end}}{{end}}

        tables = append(tables, table)
    }

    return tables, nil
}
{{end}}
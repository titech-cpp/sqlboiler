{{define "select"}}
func (q *{{.Name.UpperCamel}}Query) Find() (*{{.Name.UpperCamel}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs)
    }
    query := fmt.Sprintf("SELECT * FROM {{.Name.Snake}} LIMIT 1 %s", whereQuery)

    rows, err := q.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("Execute Query Error: %w", err)
    }
    defer rows.Close()

    table := new({{.Name.UpperCamel}}Table)
    dests := []interface{}{}{{range .Columns}}
    dests = append(dests, &(table.{{.Name.UpperCamel}})){{end}}
    if rows.Next() {
        err = rows.Scan(dests...)
        if err != nil {
            return nil, fmt.Errorf("Scaning Rows Error: %w", err)
        }
    } else {
        return nil, RECORD_NOT_FOUND
    }

    return table, nil
}

func (q *{{.Name.UpperCamel}}Query) Select() ([]*{{.Name.UpperCamel}}Table,error) {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs)
    }
    query := fmt.Sprintf("SELECT * FROM {{.Name.Snake}} %s", whereQuery)

    rows, err := q.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("Execute Query Error: %w", err)
    }
    defer rows.Close()

    tables := []*{{.Name.UpperCamel}}Table{}
    for rows.Next() {
        table := new({{.Name.UpperCamel}}Table)
        dests := []interface{}{}{{range .Columns}}
        dests = append(dests, &(table.{{.Name.UpperCamel}})){{end}}
        err = rows.Scan(dests...)
        if err != nil {
            return nil, fmt.Errorf("Scaning Rows Error: %w", err)
        }

        tables = append(tables, table)
    }

    return tables, nil
}
{{end}}
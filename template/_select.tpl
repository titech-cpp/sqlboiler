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
        return &{{.Name.UpperCamel}}Table{}, fmt.Errorf("Execute Query Error: %w", err)
    }

    table := new({{.Name.UpperCamel}}Table)
    dests := []interface{}{}{{range .Columns}}
    dests = append(dests, &(table.{{.Name.UpperCamel}})){{end}}
    rows.Next()
    err = rows.Scan(dests...)
    if err != nil {
        return &{{.Name.UpperCamel}}Table{}, fmt.Errorf("Scaning Rows Error: %w", err)
    }

    return table, nil
}
{{end}}
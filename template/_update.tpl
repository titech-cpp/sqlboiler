{{define "update"}}
func (q *{{.Name.UpperCamel}}Query) Update(t *{{.Name.UpperCamel}}) error {
    updateList := []string{}
    args := []interface{}{}{{range .Columns}}{{if not .ReadOnly}}
    if t.{{.Name.UpperCamel}} != nil {
        updateList = append(updateList, "{{.Name.Snake}}")
        args = append(args, *(t.{{.Name.UpperCamel}}))
    }{{end}}{{end}}

    if len(updateList) == 0 {
        return newError("empty update column")
    }

    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }

    cmd := fmt.Sprintf("UPDATE {{.Name.Snake}} SET %s %s", strings.Join(updateList, "=?,")+"=?", whereQuery)

    _, err := q.db.Exec(cmd, args...)
    if err != nil {
        return fmt.Errorf("failed to insert a record: %w", err)
    }

    return nil
}
{{end}}
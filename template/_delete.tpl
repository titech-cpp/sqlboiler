{{define "delete"}}
func (q *{{.Name.UpperCamel}}Query) Delete() error {
    whereMap := q.createWhereMap()
    whereQuery, whereArgs := q.whereStruct.Where(whereMap)

    args := []interface{}{}
    if len(whereArgs) != 0 {
        args = append(args, whereArgs...)
    }

    cmd := fmt.Sprintf("DELETE FROM {{.Name.Snake}} %s", whereQuery)

    _, err := q.db.Exec(cmd, args...)
    if err != nil {
        return fmt.Errorf("failed to insert a record: %w", err)
    }

    return nil
}
{{end}}
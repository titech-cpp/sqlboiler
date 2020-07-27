{{define "insert"}}
func (q *{{.Name.UpperCamel}}Query) Insert(t *{{.Name.UpperCamel}}) error {
    insertList := []string{}
    args := []interface{}{}{{range .Columns}}{{if not .ReadOnly}}
    if t.{{.Name.UpperCamel}} != nil {
        insertList = append(insertList, "{{.Name.Snake}}")
        args = append(args, *(t.{{.Name.UpperCamel}}))
    }{{end}}{{if and (not .Null) (and (not .Default) (not .AutoIncrement))}}
    if t.{{.Name.UpperCamel}} == nil {
        return newError("invalid null value")
    }{{end}}{{end}}

    if len(insertList) == 0 {
        return newError("empty insert column")
    }

    questions := []string{}
    for i := 0; i < len(args); i++ {
        questions = append(questions, "?")
    }
    cmd := fmt.Sprintf("INSERT INTO {{.Name.Snake}} (%s) VALUE (%s)", strings.Join(insertList, ","), strings.Join(questions, ","))

    _, err := q.db.Exec(cmd, args...)
    if err != nil {
        return fmt.Errorf("failed to insert a record: %w", err)
    }

    return nil
}
{{end}}
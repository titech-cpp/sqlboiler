{{define "joinedWhere"}}
func (q *Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query) createWhereMap() map[string]interface{} {
    whereMap := map[string]interface{}{}{{range .Tables}}{{$upperTable := .Name.UpperCamel}}{{range .Columns}}
    if q.table.{{printf $upperTable}}.{{.Name.UpperCamel}} != nil {
        whereMap["{{.Name.Snake}}"] = *(q.table.{{printf $upperTable}}.{{.Name.UpperCamel}})
    }{{end}}{{end}}
    return whereMap
}

func (q *Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query) Where(table *Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}) *Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query { {{range .Tables}}{{$upperTable := .Name.UpperCamel}}{{range .Columns}}
    if table.{{printf $upperTable}}.{{.Name.UpperCamel}} != nil {
        q.table.{{printf $upperTable}}.{{.Name.UpperCamel}} = table.{{printf $upperTable}}.{{.Name.UpperCamel}}
    }{{end}}{{end}}

    return q
}
{{end}}
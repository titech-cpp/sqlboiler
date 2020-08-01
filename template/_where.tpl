{{define "where"}}{{$lowerTable := .Name.LowerCamel}}{{$upperTable := .Name.UpperCamel}}
func (q *{{.Name.UpperCamel}}Query) createWhereMap() map[string]interface{} {
    whereMap := map[string]interface{}{}{{range .Columns}}
    if q.table.{{.Name.UpperCamel}} != nil {
        whereMap["{{.Name.Snake}}"] = *(q.table.{{.Name.UpperCamel}})
    }{{end}}
    return whereMap
}

func (q *{{.Name.UpperCamel}}Query) Where(table *{{.Name.UpperCamel}}) *{{.Name.UpperCamel}}Query { {{range .Columns}}
    if table.{{.Name.UpperCamel}} != nil {
        q.table.{{.Name.UpperCamel}} = table.{{.Name.UpperCamel}}
    }{{end}}

    return q
}
{{end}}
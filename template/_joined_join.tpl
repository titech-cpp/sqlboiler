{{define "joinedJoin"}}{{$v := .Tables}}{{range .Joins}}
func (q *Join{{range $v}}{{.Name.UpperCamel}}{{end}}Query) Join{{.}}() *Join{{range $v}}{{.Name.UpperCamel}}{{end}}{{.}}Query {
    return &Join{{range $v}}{{.Name.UpperCamel}}{{end}}{{.}}Query{}
}
{{end}}{{end}}
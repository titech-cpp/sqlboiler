package models
{{range .Tables}}
type {{.Name.UpperCamel}}Table struct {
{{range .Columns}}    {{.Name.UpperCamel}} {{.Type}}
{{end}}}
{{end}}
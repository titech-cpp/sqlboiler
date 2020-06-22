package models
{{range .Tables}}
type {{.Name.UpperCamel}}Table struct {
{{range .Columns}}    {{.Name.UpperCamel}} {{.Type.Code}}
{{end}}}
{{end}}
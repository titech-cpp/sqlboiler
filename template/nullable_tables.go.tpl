package models
{{range .Tables}}
type nullable{{.Name.UpperCamel}}Table struct {
{{range .Columns}}    {{.Name.UpperCamel}} {{.Type.Code.Null}}
{{end}}}
{{end}}
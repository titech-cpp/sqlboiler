package models
{{range .Tables}}
type {{.Name.UpperCamel}} struct {
{{range .Columns}}    {{.Name.UpperCamel}} *{{.Type.Code.NotNull}}
{{end}}}
{{end}}
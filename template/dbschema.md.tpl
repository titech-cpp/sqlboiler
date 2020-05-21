## {{.DB.Name}}
{{range .Tables}}
### {{.Name}}
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |{{range .Columns}}
| {{.Name}} | {{.Type}} | {{if not .Null}}No{{end}} | {{.Key}} | {{.Default}} | {{.Extra}} | {{.Description}} |{{end}}}
{{end}}
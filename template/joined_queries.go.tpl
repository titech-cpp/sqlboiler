package models

import (
    "fmt"
    "database/sql"

    "sample/models/query"
)

// Join{{range .Tables}}{{.Name.UpperCamel}}{{end}} Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}のテーブルの構造体（全ての要素がポインタ）
type Join{{range .Tables}}{{.Name.UpperCamel}}{{end}} struct { {{range .Tables}}
    {{.Name.UpperCamel}}{{end}}
}

type nullableJoin{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table struct { {{range .Tables}}
    {{.Name.UpperCamel}} nullable{{.Name.UpperCamel}}Table{{end}}
}

// Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}のテーブルの構造体
type Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Table struct { {{range .Tables}}
    {{.Name.UpperCamel}} {{.Name.UpperCamel}}Table{{end}}
}

// Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}のクエリの構造体
type Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}Query struct {
    db *sql.DB
    whereStruct *query.Where
    joins []*query.Join
    table Join{{range .Tables}}{{.Name.UpperCamel}}{{end}}
}

{{template "joinedWhere" .}}

{{template "joinedSelect" .}}

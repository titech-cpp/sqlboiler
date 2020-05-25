package model

// Schema スキーマの構造体
type Schema struct {
	DB     db
	Tables []SchemaTable
}

// SchemaTable テーブルの構造体
type SchemaTable struct {
	Name    string
	Columns []SchemaColumn
}

// SchemaColumn カラムの構造体
type SchemaColumn struct {
	Name    string
	Type    string
	Null    bool
	Key     string
	Default string
	Extra   []string
	Description string
}

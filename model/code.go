package model

// Code コード生成用の構造体
type Code struct {
	DB     db
	Tables []*CodeTable
}

// CodeTable テーブルの構造体
type CodeTable struct {
	Name    *NameDetail
	Columns []*CodeColumn
}

// CodeColumn カラムの構造体
type CodeColumn struct {
	Name     *NameDetail
	Type     string
	Null     bool
	ReadOnly bool
}

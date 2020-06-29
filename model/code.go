package model

import (
	"fmt"
)

// Code コード生成用の構造体
type Code struct {
	DB     DB
	Tables []*CodeTable
}

// Check 同一か確認
func (c *Code) Check(cd *Code) error {
	if !c.DB.Check(&cd.DB) || len(c.Tables) != len(cd.Tables) {
		return fmt.Errorf("Invalid DB or Tables Length, DB:%#v %#v, Tables:%#v %#v", c.DB, cd.DB, c.Tables, cd.Tables)
	}

	for i, v := range c.Tables {
		err := v.Check(cd.Tables[i])
		if err != nil {
			return fmt.Errorf("Invalid Table %#v %#v(Error: %w)", v, cd.Tables[i], err)
		}
	}

	return nil
}

// CodeTable テーブルの構造体
type CodeTable struct {
	Name    *NameDetail
	Columns []*CodeColumn
}

// Check 同一か確認
func (c *CodeTable) Check(ct *CodeTable) error {
	if !c.Name.Check(ct.Name) || len(c.Columns) != len(ct.Columns) {
		return fmt.Errorf("Invalid Table Name or Columns Length, Name:%#v %#v, Columns:%#v %#v", c.Name, ct.Name, c.Columns, ct.Columns)
	}

	for i, v := range c.Columns {
		if !v.Check(ct.Columns[i]) {
			return fmt.Errorf("Invalid Column %#v %#v", v, ct.Columns[i])
		}
	}

	return nil
}

// CodeColumn カラムの構造体
type CodeColumn struct {
	Name          *NameDetail
	Type          *CodeType
	Null          bool
	ReadOnly      bool
	Key           *CodeKey
	AutoIncrement bool
	Default       string
}

// Check 同一か確認
func (c *CodeColumn) Check(cc *CodeColumn) bool {
	return c.Name.Check(cc.Name) && c.Type.Check(cc.Type) && c.Null == cc.Null && c.ReadOnly == cc.ReadOnly && c.Key.Check(cc.Key) && c.AutoIncrement == cc.AutoIncrement && c.Default == cc.Default
}

// CodeType 型の構造体
type CodeType struct {
	Code string
	SQL  string
}

// Check 同一か確認
func (c *CodeType) Check(ct *CodeType) bool {
	return c.Code == ct.Code && c.SQL == ct.SQL
}

// CodeKey キーの構造体
type CodeKey struct {
	Primary bool
	Unique  bool
	Foreign []*CodeForeign
}

// Check 同一か確認
func (c *CodeKey) Check(ck *CodeKey) bool {
	for i, v := range c.Foreign {
		if !ck.Foreign[i].Check(v) {
			return false
		}
	}

	return c.Primary == ck.Primary && c.Unique == ck.Unique && len(c.Foreign) == len(ck.Foreign)
}

// CodeForeign 外部キーの構造体
type CodeForeign struct {
	Table  string
	Column string
}

// Check 同一か確認
func (c *CodeForeign) Check(cf *CodeForeign) bool {
	return c.Table == cf.Table && c.Column == cf.Column
}

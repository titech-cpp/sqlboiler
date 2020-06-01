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
func (c *Code)Check(cd *Code) error {
	if !c.DB.Check(&cd.DB) || len(c.Tables) != len(cd.Tables) {
		return fmt.Errorf("Invalid DB or Tables Length, DB:%#v %#v, Tables:%#v %#v", c.DB, cd.DB, c.Tables, cd.Tables)
	}

	for i,v := range c.Tables {
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
func (c *CodeTable)Check(ct *CodeTable) error {
	if !c.Name.Check(ct.Name) || len(c.Columns) != len(ct.Columns) {
		return fmt.Errorf("Invalid Table Name or Columns Length, Name:%#v %#v, Columns:%#v %#v", c.Name, ct.Name, c.Columns, ct.Columns)
	}

	for i,v := range c.Columns {
		if !v.Check(ct.Columns[i]) {
			return fmt.Errorf("Invalid Column %#v %#v", v, ct.Columns[i])
		}
	}

	return nil
}

// CodeColumn カラムの構造体
type CodeColumn struct {
	Name     *NameDetail
	Type     string
	Null     bool
	ReadOnly bool
}

// Check 同一か確認
func (c *CodeColumn) Check(cc *CodeColumn) bool {
	return c.Name.Check(cc.Name) && c.Type == cc.Type && c.Null == cc.Null && c.ReadOnly == cc.ReadOnly
}

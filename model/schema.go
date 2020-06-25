package model

// Schema スキーマの構造体
type Schema struct {
	DB     DB
	Tables []*SchemaTable
}

// Check 同一か確認
func (s *Schema) Check(sc *Schema) bool {
	if !s.DB.Check(&sc.DB) || len(s.Tables) != len(sc.Tables) {
		return false
	}

	for i, v := range s.Tables {
		if !v.Check(sc.Tables[i]) {
			return false
		}
	}

	return true
}

// SchemaTable テーブルの構造体
type SchemaTable struct {
	Name    string
	Columns []*SchemaColumn
}

// Check 同一か確認
func (s *SchemaTable) Check(st *SchemaTable) bool {
	if s.Name != st.Name || len(s.Columns) != len(st.Columns) {
		return false
	}

	for i, v := range s.Columns {
		if !v.Check(st.Columns[i]) {
			return false
		}
	}

	return true
}

// SchemaColumn カラムの構造体
type SchemaColumn struct {
	Name        string
	Type        string
	Null        bool
	Key        []string
	Default     string
	Extra       []string
	Description string
}

// Check 同一か確認
func (s *SchemaColumn) Check(sc *SchemaColumn) bool {
	for i,v := range s.Key {
		if sc.Key[i] != v {
			return false
		}
	}

	for i, v := range s.Extra {
		if sc.Extra[i] != v {
			return false
		}
	}

	return s.Name == sc.Name && s.Type == sc.Type && s.Null == sc.Null && len(s.Key) == len(sc.Key) && s.Default == sc.Default && len(s.Extra) == len(sc.Extra) && s.Description == sc.Description
}

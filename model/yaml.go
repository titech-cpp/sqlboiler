package model

// Yaml Yamlの構造体
type Yaml struct {
	DB     DB
	Tables map[string]yamlColumns
}

// Check 同一か確認
func (y *Yaml) Check(yml *Yaml) bool {
	if !y.DB.Check(&yml.DB) {
		return false
	}

	for key, value := range y.Tables {
		table, ok := yml.Tables[key]
		if !ok {
			return false
		}
		for k, val := range value {
			v, ok := table[k]
			if !ok || !val.Check(v) {
				return false
			}
		}
		for k, val := range table {
			v, ok := value[k]
			if !ok || !val.Check(v) {
				return false
			}
		}
	}

	for key, value := range yml.Tables {
		table, ok := y.Tables[key]
		if !ok {
			return false
		}
		for k, val := range value {
			v, ok := table[k]
			if !ok || !val.Check(v) {
				return false
			}
		}
		for k, val := range table {
			v, ok := value[k]
			if !ok || !val.Check(v) {
				return false
			}
		}
	}

	return true
}

type yamlColumns = map[string]*YamlColumn

// YamlColumn カラムの構造体
type YamlColumn struct {
	Type          string
	Null          bool
	AutoIncrement bool
	Key           string
	Default       string
}

// Check 同一か確認
func (y *YamlColumn) Check(yc *YamlColumn) bool {
	return y.Type == yc.Type && y.Null == yc.Null && y.AutoIncrement == yc.AutoIncrement && y.Key == yc.Key && y.Default == yc.Default
}

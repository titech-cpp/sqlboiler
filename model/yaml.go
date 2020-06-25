package model

// Yaml Yamlの構造体
type Yaml struct {
	DB     DB
	Tables map[string]yamlColumns
}

// Check 同一か確認
func (y *Yaml) Check(yml *Yaml) bool {
	for key, value := range y.Tables {
		table, ok := yml.Tables[key]
		if !ok {
			return false
		}
		for i, val := range table {
			if !val.Check(value[i]) {
				return false
			}
		}
	}

	return y.DB.Check(&yml.DB) && len(y.Tables) == len(yml.Tables)
}

type yamlColumns = []*YamlColumn

// YamlColumn カラムの構造体
type YamlColumn struct {
	Name          string
	Type          string
	NoNull        bool `yaml:"no_null"`
	AutoIncrement bool `yaml:"auto_increment"`
	Key           string
	Default       string
}

// Check 同一か確認
func (y *YamlColumn) Check(yc *YamlColumn) bool {
	return y.Name == yc.Name || y.Type == yc.Type && y.NoNull == yc.NoNull && y.AutoIncrement == yc.AutoIncrement && y.Key == yc.Key && y.Default == yc.Default
}

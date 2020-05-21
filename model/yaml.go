package model

// Yaml Yamlを読み取った後の構造体
type Yaml struct {
	DB     db
	Tables map[string]yamlColumns
}

type yamlColumns = map[string]yamlColumn

type yamlColumn struct {
	Type          string
	Null          bool
	AutoIncrement bool
	Key           string
	Default       string
}

package query

import (
	"fmt"
)

const tagName = "db"

// Where WHERE関連の構造体
type Where struct {}

// Where WHEREのクエリを作る
func (*Where) Where(whereMap map[string]interface{}) (string, []interface{}) {
	where := ""
	args := make([]interface{}, 0)
	if len(whereMap) != 0 {
		where = "WHERE"
	}
	for key, val := range whereMap {
		where += fmt.Sprintf(" %s = ?", key)
		args = append(args, val)
	}
	return where, args
}

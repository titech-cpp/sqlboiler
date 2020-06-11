package query

import (
	"fmt"
)

// Where WHERE関連の構造体
type Where struct {}

// Where WHEREのクエリを作る
func (*Where) Where(whereMap map[string]interface{}) (string, []interface{}) {
	where := ""
	args := make([]interface{}, 0)
	if len(whereMap) != 0 {
		where = "WHERE"
	}

	i := 0
	for key, val := range whereMap {
		i++
		where += fmt.Sprintf(" %s = $%d", key, i)
		args = append(args, val)
	}

	return where, args
}

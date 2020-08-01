package query

import "fmt"

//JoinDirection JOINの向き
type JoinDirection int

//JoinType JOINの種類
type JoinType int

const (
	LEFT JoinDirection = iota
	RIGHT
	INNER JoinType = iota
	OUTER
)

// Join Joinの状況
type Join []string

// NewJoin Joinのコンストラクタ
func NewJoin(joinType JoinType, firstKey string, secondKey string, directions ...JoinDirection) *Join {
	switch joinType {
	case INNER:
		return &Join{"INNER JOIN", fmt.Sprintf("ON %s = %s", firstKey, secondKey)}
	case OUTER:
		switch directions[0] {
		case LEFT:
			return &Join{"LEFT OUTER JOIN", fmt.Sprintf("ON %s = %s", firstKey, secondKey)}
		case RIGHT:
			return &Join{"RIGHT OUTER JOIN", fmt.Sprintf("ON %s = %s", firstKey, secondKey)}
		}
	}

	return nil
}

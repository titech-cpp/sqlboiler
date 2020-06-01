package model

import (
	"fmt"
	"strings"
)

// DB データベース
type DB struct {
	Type string
	Name string
}

// Check 同一か確認
func (d *DB) Check(db *DB) bool {
	return d.Type == db.Type && d.Name == db.Name
}

// NameDetail 名前の詳細
type NameDetail struct {
	UpperCamel string
	LowerCamel string
	Snake      string
}

// NewNameDetail NameDeatailのコンストラクタ
func NewNameDetail(snake string) (*NameDetail, error) {
	nameDetail := new(NameDetail)
	nameDetail.Snake = snake
	upper, err := nameDetail.snakeToCamel(snake, true)
	if err != nil {
		return nil, fmt.Errorf("Parse Snake Case To Upper Camel Case Error: %w", err)
	}
	nameDetail.UpperCamel = upper
	lower, err := nameDetail.snakeToCamel(snake, false)
	if err != nil {
		return nil, fmt.Errorf("Parse Snake Case To Lower Case Error: %w", err)
	}
	nameDetail.LowerCamel = lower
	return nameDetail, nil
}

// Check 同一かチェック
func (n *NameDetail) Check(nd *NameDetail) bool {
	return n.UpperCamel == nd.UpperCamel && n.LowerCamel == nd.LowerCamel && n.Snake == nd.Snake
}

func (*NameDetail) snakeToCamel(snake string, isUpper bool) (string, error) {
	var builder strings.Builder
	isUnderBar := false
	for i, c := range snake {
		if (c < 'a' || 'z' < c) && c != '_' {
			return "", fmt.Errorf("Invaid Non Small Case %s In %s", string(c), snake)
		}
		if i == 0 && c == '_' {
			return "", fmt.Errorf("Invalid First `_` In %s", snake)
		} else if i == 0 && isUpper {
			builder.WriteRune(c - 'a' + 'A')
		} else if isUnderBar && c == '_' {
			return "", fmt.Errorf("Invalid Consecutive UnderBar In %s", snake)
		} else if isUnderBar {
			builder.WriteRune(c - 'a' + 'A')
			isUnderBar = false
		} else if c == '_' {
			isUnderBar = true
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String(), nil
}

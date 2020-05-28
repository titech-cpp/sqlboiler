package model

import (
	"bytes"
	"fmt"
)

type db struct {
	Type string
	Name string
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
	upper, err := snakeToCamel(snake, true)
	if err != nil {
		return nil, fmt.Errorf("Parse Snake Case To Upper Camel Case Error: %w", err)
	}
	nameDetail.UpperCamel = upper
	lower, err := snakeToCamel(snake, false)
	if err != nil {
		return nil, fmt.Errorf("Parse Snake Case To Lower Case Error: %w", err)
	}
	nameDetail.LowerCamel = lower
	return nameDetail, nil
}

func snakeToCamel(snake string, isUpper bool) (string, error) {
	var buf bytes.Buffer
	isUnderBar := false
	for i, c := range snake {
		if (c < 'a' || 'z' < c) && c != '_' {
			return "", fmt.Errorf("Invaid Non Small Case %s In %s", string(c), snake)
		}
		if i == 0 && c == '_' {
			return "", fmt.Errorf("Invalid First `_` In %s", snake)
		} else if i == 0 && isUpper {
			buf.WriteRune(c - 'a' + 'A')
		} else if isUnderBar && c == '_' {
			return "", fmt.Errorf("Invalid Consecutive UnderBar In %s", snake)
		} else if isUnderBar {
			buf.WriteRune(c - 'a' + 'A')
			isUnderBar = false
		} else if c == '_' {
			isUnderBar = true
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String(), nil
}

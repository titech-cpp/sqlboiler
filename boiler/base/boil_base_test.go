package base

import (
	"io"
	"strings"
	"testing"

	"github.com/titech-cpp/sqlboiler/model"
)

func TestBoilBase(t *testing.T) {
	boilBase := NewBoilBase("test")
	nameDetail, err := model.NewNameDetail("test")
	if err != nil {
		t.Fatalf("Name Detail Constructor Error: %#v", err)
	}
	code := &model.Code{
		DB: model.DB{
			Type: "mysql",
			Name: "test-db",
		},
		Tables: []*model.CodeTable{
			{
				Name: nameDetail,
				Columns: []*model.CodeColumn{
					{
						Name: nameDetail,
						Type: &model.CodeType{
							Code: "test",
							SQL: "test",
						},
						Null: true,
						ReadOnly: true,
					},
				},
			},
		},
	}
	expect := `package models

type TestTable struct {
    Test test
}
`
	pr,pw := io.Pipe()

	errCh :=make(chan error)
	go func() {
		err = boilBase.MakeFile(pw, "tables.go", code)
		errCh <- err
	}()

	var builder strings.Builder
	_,err = io.Copy(&builder, pr)
	if err != nil {
		t.Fatalf("Unexpected Copy Error: %#v", err)
	}

	err = <-errCh
	if err != nil {
		t.Fatalf("Unexpected MakeFile Error: %#v", err)
	}

	str := builder.String()
	if str != expect {
		t.Fatalf("Invalid Created File Value %s, Expected %s", str, expect)
	}
}
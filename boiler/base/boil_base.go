package base

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/rakyll/statik/fs"
	_ "github.com/titech-cpp/sqlboiler/statik"
)

// BoilerBase ボイラーのインターフェイス
type BoilerBase interface {
	MakeFile(string, interface{}) error
}

// BoilBase ファイル生成の構造体
type BoilBase struct {
	basePath string
}

// NewBoilBase Boilのコンストラクタ
func NewBoilBase(basePath string) *BoilBase {
	boil := BoilBase{
		basePath: basePath,
	}
	return &boil
}

// MakeFile ファイルの生成
func (b *BoilBase) MakeFile(fileName string, vars interface{}) error {
	FS, err := fs.New()
	if err != nil {
		return fmt.Errorf("Statik New Error: %w", err)
	}

	fp, err := FS.Open("/" + fileName + ".tpl")
	if err != nil {
		return fmt.Errorf("Statik Open Error: %w", err)
	}
	defer fp.Close()

	file, err := os.Create(b.basePath + "/" + fileName)
	if err != nil {
		return fmt.Errorf("File Create Error: %w", err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(fp)
	if err != nil {
		return fmt.Errorf("Read File Error: %w", err)
	}
	tmpFile := buf.String()

	tmp, err := template.New(fileName).Parse(tmpFile)
	if err != nil {
		return fmt.Errorf("Parse Template Error: %w", err)
	}

	err = tmp.Execute(file, vars)
	if err != nil {
		return fmt.Errorf("Execute Template Error: %w", err)
	}

	return nil
}

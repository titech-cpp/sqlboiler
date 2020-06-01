package base

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/rakyll/statik/fs"
	_ "github.com/titech-cpp/sqlboiler/statik"
)

// BoilerBase ボイラーのインターフェイス
type BoilerBase interface {
	MakeBaseDir() error
	MakeFile(io.WriteCloser, string, interface{}) error
	MakeFileWriter(string) (io.WriteCloser, error)
}

// BoilBase ファイル生成の構造体
type BoilBase struct {
	basePath string
}

// NewBoilBase Boilのコンストラクタ
func NewBoilBase(basePath string) *BoilBase {
	boil := new(BoilBase)
	boil.basePath = basePath
	return boil
}

// MakeBaseDir ファイル作成先のディレクトリの作成
func (b * BoilBase) MakeBaseDir() error {
	err := os.MkdirAll(b.basePath, 0777)
	if err != nil {
		return err
	}

	return nil
}

// MakeFile ファイルの生成
func (b *BoilBase) MakeFile(file io.WriteCloser, tmpNameBase string, vars interface{}) error {
	defer file.Close()

	FS, err := fs.New()
	if err != nil {
		return fmt.Errorf("Statik New Error: %w", err)
	}

	fp, err := FS.Open("/" + tmpNameBase + ".tpl")
	if err != nil {
		return fmt.Errorf("Statik Open Error: %w", err)
	}
	defer fp.Close()

	var builder strings.Builder
	_, err = io.Copy(&builder, fp)
	if err != nil {
		return fmt.Errorf("Read File Error: %w", err)
	}
	tmpFile := builder.String()

	tmp, err := template.New(tmpNameBase).Parse(tmpFile)
	if err != nil {
		return fmt.Errorf("Parse Template Error: %w", err)
	}

	err = tmp.Execute(file, vars)
	if err != nil {
		return fmt.Errorf("Execute Template Error: %w", err)
	}

	return nil
}

// MakeFileWriter ファイルのWriterの作成
func (b *BoilBase) MakeFileWriter(fileName string) (io.WriteCloser, error) {
	file, err := os.Create(b.basePath + "/" + fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

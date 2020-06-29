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
func (b *BoilBase) MakeBaseDir() error {
	err := os.RemoveAll(b.basePath)
	if err != nil {
		return fmt.Errorf("Remove Files Error: %w", err)
	}
	err = os.MkdirAll(b.basePath, 0777)
	if err != nil {
		return fmt.Errorf("Make Directory Error: %w", err)
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

	defines := []string{"select"}
	defineFiles := make([]string, 0, len(defines))
	for _, v := range defines {
		fp, err = FS.Open("/_" + v + ".tpl")
		if err != nil {
			return fmt.Errorf("Statik Open Error: %w", err)
		}
		defer fp.Close()

		var builder strings.Builder
		_, err = io.Copy(&builder, fp)
		if err != nil {
			return fmt.Errorf("Read File Error(%s): %w", v, err)
		}
		file := builder.String()
		defineFiles = append(defineFiles, file)
	}

	funcMap := map[string]interface{}{
		"isLast": func(i int, l int) bool {
			return i == l-1
		},
	}

	tmp, err := template.New(tmpNameBase).Funcs(funcMap).Parse(tmpFile)
	if err != nil {
		return fmt.Errorf("Parse Template Error: %w", err)
	}
	for i, v := range defineFiles {
		tmp, err = tmp.Parse(v)
		if err != nil {
			return fmt.Errorf("Parse Template Error(%s): %w", defines[i], err)
		}
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

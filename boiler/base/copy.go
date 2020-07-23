package base

import (
	"fmt"
	"io"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/titech-cpp/sqlboiler/statik"
)

// CopyBase ボイラーのインターフェイス
type CopyerBase interface {
	MakeBaseDir() error
	MakeFile(io.WriteCloser, string) error
	MakeFileWriter(string) (io.WriteCloser, error)
}

// CopyBase ファイル生成の構造体
type CopyBase struct {
	basePath string
}

// NewCopyBase Boilのコンストラクタ
func NewCopyBase(basePath string) *CopyBase {
	copy := new(CopyBase)
	copy.basePath = basePath
	return copy
}

// MakeBaseDir ファイル作成先のディレクトリの作成
func (b *CopyBase) MakeBaseDir() error {
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
func (b *CopyBase) MakeFile(file io.WriteCloser, nameBase string) error {
	defer file.Close()

	FS, err := fs.New()
	if err != nil {
		return fmt.Errorf("Statik New Error: %w", err)
	}

	fp, err := FS.Open("/" + nameBase + ".go")
	if err != nil {
		return fmt.Errorf("Statik Open Error: %w", err)
	}
	defer fp.Close()

	defer file.Close()

	_, err = io.Copy(file, fp)
	if err != nil {
		return fmt.Errorf("File Copy Error: %w", err)
	}

	return nil
}

// MakeFileWriter ファイルのWriterの作成
func (b *CopyBase) MakeFileWriter(fileName string) (io.WriteCloser, error) {
	file, err := os.Create(b.basePath + "/" + fileName + ".go")
	if err != nil {
		return nil, err
	}

	return file, nil
}

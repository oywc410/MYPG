package backup

import (
	"os"
	"path/filepath"
	"archive/zip"
	"io"
)

type Archiver interface {
	DestFmt() string
	Archive(src, dest string) error
}

type zipper struct {
}

//外部引用时可以直接暴露Archive函数
var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	//walk遍历所有目录(包括子目录)
	//加入到zip包
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) {
		if info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		io.Copy(f, in)
		return nil
	})
}

func (z *zipper) DestFmt() string {
	return "%d.zip"
}
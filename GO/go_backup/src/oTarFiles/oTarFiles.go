package oTarFiles

import (
	"archive/tar"
	"io"
	getAllFiles "oGetAllFiles"
	"os"
	"path"
)

type tarObj struct {
	fw io.Writer
	tw *tar.Writer
}

func CreateTar(getWriter func() io.Writer) *tarObj {
	return &tarObj{fw: getWriter(), tw: nil}
}

func (t *tarObj) GetTarFw() io.Writer {
	return t.fw
}

func (t *tarObj) GetTarTw() *tar.Writer {
	return t.tw
}

func (t *tarObj) AddFile(filePath, catPath string) {
	filePath = path.Clean(filePath)
	catPath = path.Clean(catPath)
	// 获取文件或目录信息
	fi, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}

	t.add(fi, filePath, catPath)
}

func (t *tarObj) AddFileInfoSt(fileInfoSts getAllFiles.FileInfoSt, catPath string) {
	catPath = path.Clean(catPath)
	t.add(fileInfoSts.GetFileInfo(), fileInfoSts.GetFilePath(), catPath)
}

func (t *tarObj) add(fileInfo os.FileInfo, filePath, catPath string) {
	t.newTarFile()
	if fileInfo.IsDir() {
		catPath += string(os.PathSeparator)
	}

	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		panic(err)
	}
	hdr.Name = catPath

	if err = t.tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	//写文件
	if !fileInfo.IsDir() {

		//打开文件
		fr, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		_, err = io.Copy(t.tw, fr)
		if err != nil {
			panic(err)
		}
	}
}

func (t *tarObj) newTarFile() {

	if t.tw == nil {
		t.tw = tar.NewWriter(t.fw)
	}
}

func (t *tarObj) CloseFile() {
	if t.tw != nil {
		err := t.tw.Close()
		if err != nil {
			panic(err)
		}
	}
}

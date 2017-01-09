package getAllFiles

import (
	"os"
)

//文件信息
type FileInfoSt struct {
	FilePath string
	FileInfo os.FileInfo
}

func (fileInfo FileInfoSt) GetFilePath() string {
	return fileInfo.FilePath
}

func (fileInfo FileInfoSt) GetFileInfo() os.FileInfo {
	return fileInfo.FileInfo
}

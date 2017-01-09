package oUnTarFiles

type FileChan struct {
	fileByte []byte
	err      error
}

func (fileChan FileChan) GetFileByte() []byte {
	return fileChan.fileByte
}

func (fileChan FileChan) GetFileErr() error {
	return fileChan.err
}

package backup

import (
	"crypto/md5"
	"path/filepath"
	"os"
	"io"
	"fmt"
)

func DirHash(path string) (string, error) {
	hash := md5.New()
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		io.WriteString(hash, path)
		fmt.Sprintf(hash, "%v", info.IsDir())
		fmt.Sprintf(hash, "%v", info.ModTime())
		fmt.Sprintf(hash, "%v", info.Mode())
		fmt.Sprintf(hash, "%v", info.Name())
		fmt.Sprintf(hash, "%v", info.Size())
		return nil
	})

	if err != nil {
		return "", err
	}
	//%x 16进制
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

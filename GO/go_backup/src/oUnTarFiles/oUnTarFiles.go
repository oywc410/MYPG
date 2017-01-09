package oUnTarFiles

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

type unTarFiles struct {
	fw io.Reader
	tr *tar.Reader
}

func CreateUnTar(tarFw io.Reader) *unTarFiles {
	return &unTarFiles{fw: tarFw, tr: nil}
}

func (obj *unTarFiles) GetAllFileChan() {
	//读取解压缩文件
	obj.tr = tar.NewReader(obj.fw)
	i := 0
	for {
		i++
		//获取压缩文件头
		hdr, er := obj.tr.Next()

		if er == io.EOF {
			break
		} else if er != nil {
			panic(er)
		}
		//文件信息
		fi := hdr.FileInfo()

		if i%1000 == 0 {
			fmt.Println(hdr.Name)
		}

		path := "tmp/" + hdr.Name

		if hdr.Typeflag == tar.TypeDir {
			// 创建目录
			os.MkdirAll(path, fi.Mode().Perm())
			// 设置目录权限
			os.Chmod(path, fi.Mode().Perm())

		} else {
			// 创建文件所在的目录
			os.MkdirAll(catNewDir(path), os.ModePerm)
			os.Chmod(catNewDir(path), fi.Mode().Perm())

			//创建文件
			f, err := os.Create(path)
			if err != nil {
				panic(err)
			}

			//获取文件流
			fileChan := obj.getFileChan()
		LOOP:
			for {
				select {
				case fileByteData := <-fileChan:
					if fileByteData.err == io.EOF {
						//文件流读取完毕
						break LOOP
					} else if fileByteData.err != nil {
						panic(fileByteData.err)
					}
					//写入文件
					f.Write(fileByteData.GetFileByte())

				}
			}

			// 设置文件权限
			os.Chmod(path, fi.Mode().Perm())
		}
	}
}

func (obj *unTarFiles) getFileChan() <-chan FileChan {
	chanFileByteData := make(chan FileChan, 50)
	go func() {
		for {
			//数据流容器
			fileChan := FileChan{fileByte: make([]byte, 1024), err: nil}

			num, err := io.ReadFull(obj.tr, fileChan.fileByte)

			//读取的数据小于1024大小时返回ErrUnexpectedEOF错误
			if err == io.ErrUnexpectedEOF {
				fileChan.fileByte = fileChan.fileByte[:num]
			} else {
				fileChan.err = err
			}

			if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
				panic(err)
			}
			//数据流写入通道中
			chanFileByteData <- fileChan

			if err == io.EOF {
				break
			}
		}
	}()

	return chanFileByteData
}

func catNewDir(str string) string {
	rs := []rune(str)
	end := 0
	for keys, a := range rs {
		if a == 92 {
			end = keys
		}
	}

	return string(rs[:end])
}

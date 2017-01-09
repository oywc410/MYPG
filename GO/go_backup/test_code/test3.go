package main

import (
	"fmt"
	"io"
	"os"
)

type fileRe struct {
	fileByte []byte
	err      error
}

//指定大小读取文件操作
func main() {
	r, _ := os.Open("go_tar_gzip4.go")

	fileReChan := make(chan fileRe, 10)

	go func() {

		for {
			obj := fileRe{fileByte: make([]byte, 1024), err: nil}

			num, err := io.ReadFull(r, obj.fileByte)
			defer r.Close()
			obj.err = err

			//读取的数据小于1024时
			if err == io.ErrUnexpectedEOF {
				obj.fileByte = obj.fileByte[:num]
			}

			fileReChan <- obj

			if err == io.EOF {
				break
			}
		}
	}()

LOOP:
	for {
		select {
		case test := <-fileReChan:
			if test.err == io.EOF {
				break LOOP
			}

			fmt.Println(string(test.fileByte))

		}
	}

}

//压缩文件
package main

//go get golang.org/x/text

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

var pathChan chan string

func main() {

	pathChan = make(chan string, 10)

	runtime.GOMAXPROCS(3)

	t1 := time.Now()

	tarFile("C:/WEB/www/")
	//tarFile("C:/MYPG/go_backup/file/")

	//main1("C:/WEB/www/")
	fmt.Println(time.Now().Sub(t1))
}

func getAllFile(dirPath string) <-chan string {

	//pathChan := make(chan string, 10)

	go func(dirPath string) {
		//打开文件夹
		dir, err := os.Open(dirPath)
		if err != nil {
			panic(err)
		}
		defer dir.Close()

		//读取文件列表
		fis, err := dir.Readdir(0) //n<=0 读取所有 n>0 读取n个
		if err != nil {
			panic(err)
		}

		var arrDirs []string

		for _, fi := range fis {
			if fi.IsDir() {
				arrDirs = append(arrDirs, dirPath+fi.Name()+"/")
			} else {
				runtime.Gosched()
				pathChan <- dirPath + fi.Name()
			}
		}

		lens := len(arrDirs)
		if lens > 0 {

			var wg sync.WaitGroup

			wg.Add(lens)

			for i := 0; i < lens; i++ {
				go func(pathDir string) {
					nextPathChange := getAllFile(pathDir)
				LOOP:
					for {
						select {
						case pathString := <-nextPathChange:
							if pathString == "!END!" {
								break LOOP
							} else {
								pathChan <- pathString
							}
						}
					}
					wg.Done()
				}(arrDirs[i])
			}

			wg.Wait()
		}

		pathChan <- "!END!"
	}(dirPath)

	return pathChan

}

func tarFile(dirPath string) {

	//建立写文件
	fw, err := os.Create("tar/lin_golang_src.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	//gizp write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	//tar write
	tw := tar.NewWriter(gw)

	defer tw.Close()

	addAllFile(dirPath, tw)

	fmt.Println("tar.gz ok")

}

/**
字符串截取
*/
func catPath(str string, start int) string {
	rs := []rune(str)
	return string(rs[start:])
}

/**
添加文件至tar
*/
func addAllFile(dirPath string, tw *tar.Writer) {

	//len := len(dirPath)
	//fmt.Println(len)
	//fmt.Println(catPath(dirPath+"asdsad/asdsa/asd.sad", len))

	i := 0

	pathChan := getAllFile(dirPath)

LOOP:
	for {
		select {
		case pathString := <-pathChan:
			if pathString == "!END!" {
				break LOOP
			} else {
				//fmt.Println(pathString)
				i++

				func(tw *tar.Writer) {
					fileInfo, err := os.Stat(pathString)
					if err != nil {
						panic(err)
					}

					//打开文件
					fr, err := os.Open(pathString)

					if err != nil {
						panic(err)
					}
					defer fr.Close()

					hdr, err := tar.FileInfoHeader(fileInfo, "")

					hdr.Name = fr.Name()
					if err = tw.WriteHeader(hdr); err != nil {
						panic(err)
					}

					//写文件
					_, err = io.Copy(tw, fr)
					if err != nil {
						panic(err)
					}
				}(tw)

			}
		}
	}

	fmt.Println(i)
}

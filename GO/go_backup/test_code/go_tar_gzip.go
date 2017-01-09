//压缩文件
package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	t1 := time.Now()
	main1()
	fmt.Println(time.Now().Sub(t1))
}

func main1() {
	//timeout := time.After(time.Second) //时间计时器(指针)

	allFilePath := getAllFile("file/")

	i := 0

LOOP: //名称任意
	for {
		select {
		case filepath := <-allFilePath:
			i++
			fmt.Println(filepath)
			if filepath == "!allFileEnd" {
				fmt.Println(i)
				break LOOP
			}
			//fmt.Println(i)
		default:
			//fmt.Println("-")
		}
	}
}

func main2() {
	getAllFile2("file/")
	fmt.Println(j)
}

var j int

func getAllFile2(dirPath string) {
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

	for _, fi := range fis {
		if fi.IsDir() {
			getAllFile2(dirPath + fi.Name() + "/")
		} else {
			fmt.Println(dirPath + fi.Name())
			j++
		}
	}
}

var empty struct{} // 限制大小用

func getAllFile(dirPath string) <-chan string {

	//固定处理数
	filePathChan := make(chan string, 50)

	//固定最大获取数
	limit := make(chan struct{}, 100)

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

	go func(fis []os.FileInfo, dir *os.File, dirPath string) {

		for _, fi := range fis {
			select {
			case limit <- empty:
				go func(fi os.FileInfo, dir *os.File, dirPath string) {
					if fi.IsDir() {
						go func(fi os.FileInfo, dir *os.File, dirPath string) {
							allFilePath := getAllFile(dirPath + fi.Name() + "/")
						LOOP: //名称任意
							for {
								select {
								case filepath := <-allFilePath:
									//fmt.Println(filepath)
									if filepath == "!allFileEnd" {
										break LOOP
									} else {
										filePathChan <- filepath
									}
								}
							}
						}(fi, dir, dirPath)

					} else {
						//打印文件名称
						//fmt.Println(dirPath + fi.Name())
						//获取文件名
						filePathChan <- dir.Name() + fi.Name()
					}

					<-limit
				}(fi, dir, dirPath)
			}
		}

		//filePathChan <- "!allFileEnd"

	}(fis, dir, dirPath)

	/*
		//遍历文件列表
		for _, fi := range fis {
			if fi.IsDir() {
				addAllFile(dirPath+fi.Name()+"/", tw)
				continue
			}

			//打印文件名称
			fmt.Println(dirPath + fi.Name())
			//获取文件名

			fielPath := dir.Name() + "/" + fi.Name()
		}
	*/

	return filePathChan
}

func tarFile() {

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

	addAllFile("file/", tw)

	fmt.Println("tar.gz ok")

}

func addAllFile(dirPath string, tw *tar.Writer) {
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

	//遍历文件列表
	for _, fi := range fis {
		if fi.IsDir() {
			addAllFile(dirPath+fi.Name()+"/", tw)
			continue
		}

		//打印文件名称
		fmt.Println(dirPath + fi.Name())

		//打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		//信息头
		h := new(tar.Header)
		h.Name = dirPath + fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()

		//写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}

		//写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
}

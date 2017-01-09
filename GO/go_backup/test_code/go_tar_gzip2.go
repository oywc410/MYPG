//压缩文件
package main

import (
	"fmt"
	//"io"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)

	t1 := time.Now()
	main2("C:/WEB/www/")
	fmt.Println(time.Now().Sub(t1))
}

func main1(path string) {
	getAllFile1(path)
	fmt.Println(j)
}

func main2(path string) {
	getAllFile2(path)
	fmt.Println(j)
}

func getAllFile1(dirPath string) {

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
			//fmt.Println(dirPath + fi.Name())
			j++
		}
	}

	lens := len(arrDirs)
	if lens > 0 {

		var wg sync.WaitGroup

		wg.Add(lens)

		for i := 0; i < lens; i++ {
			go func(pathDir string) {
				getAllFile1(pathDir)
				wg.Done()
			}(arrDirs[i])
		}

		wg.Wait()
	}

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
			//fmt.Println(dirPath + fi.Name())
			j++

			if j%1000 == 0 {
				fmt.Println(j)
				fmt.Println(dirPath + fi.Name())
			}

		}
	}

}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3) //增加  (开启阻塞)
	go func() {
		downloadNodeVersion("0.4.7")
		wg.Done() //减少			2
	}()

	go func() {
		downloadNodeVersion("0.4.8")
		wg.Done() //减少			2
	}()

	go func() {
		downloadNodeVersion("0.4.9")
		wg.Done() //减少			2
	}()

	wg.Wait() //阻塞结束  为0时结束阻塞
	fmt.Println("下载完成")
}

//异步下载
func downloadNodeVersion(v string) {
	fmt.Println("下载" + v)
	c := exec.Command("G:/Program Files/Git/mingw64/bin/curl.exe", "http://nodejs.org/dist/node-v"+v+".tar.gz")

	reData, err := c.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		file, err := os.Create("./tmp/node-v" + v + ".tar.gz")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		//写入文件
		_, err = file.Write(reData)
	}

}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

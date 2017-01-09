package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件夹
	os.Mkdir("aaa", 0777)
	//创建连续的子文件夹
	os.MkdirAll("aaa/test1/test2", 0777)
	//删除单个文件夹
	err := os.Remove("aaa")
	if err != nil {
		fmt.Println(err)
	}
	//删除子文件夹(aaa不会被删除)
	os.RemoveAll("aaa")
}

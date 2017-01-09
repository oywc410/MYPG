package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//字符串读取对象
	strReader := strings.NewReader("hello world")

	//放入缓存对象
	bufReader := bufio.NewReader(strReader)

	//指定读取的缓存大小,读取值不会删除
	data, _ := bufReader.Peek(5)

	//bufReader.Buffered()缓存的字符数 (11)
	fmt.Println(string(data), bufReader.Buffered())

	//切割读取(读取到' '为止),读取值会被删除 bufReader.Buffered() 5
	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())

	//屏幕
	w := bufio.NewWriter(os.Stdout)

	//写入文件
	fmt.Fprint(w, "Hello ")
	fmt.Fprint(w, "world!\n")
	//将缓存提交至屏幕
	w.Flush()
}

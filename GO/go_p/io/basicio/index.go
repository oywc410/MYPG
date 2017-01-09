package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {

	p := make([]byte, num)
	//n读取数量
	n, err := reader.Read(p)

	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

//字符串读取
func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Print(string(data))
}

//输入os读取
func sampleReadStdin() {
	fmt.Println("please input from stdin:")

	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

//文件读取
func sampleReadFile() {
	file, _ := os.Open("index.go")
	defer file.Close()

	data, _ := ReadFrom(file, 9)

	fmt.Println(data)
}

func main() {
	//sampleReadFromString()
	//sampleReadStdin()
	sampleReadFile()
}

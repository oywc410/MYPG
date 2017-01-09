package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	Create()
	Open()
	ReadAll()
	WriteAll()
}

//创建文件
func Create() {
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//写入文件
	message := []byte("hello world\n")
	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	//直接写入文字
	_, err = file.WriteString("hello world\n")

	_, err = fmt.Fprint(file, "hello world\n")
}

//打开文件
func Open() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	message := make([]byte, 12)
	_, err = file.Read(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(message))
	_, err = file.Read(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(message))
}

//读取所有信息
func ReadAll() {
	file, _ := os.Open("./file.txt")
	message, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))
	defer file.Close()
	//message, err := ioutil.ReadFile("./file.txt")
}

//写入所有信息
func WriteAll() {

	message := []byte("hello world\n")
	err := ioutil.WriteFile("./file.txt", message, 0666)
	if err != nil {
		log.Fatal(err)
	}

}

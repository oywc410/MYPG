package main

import (
	"fmt"
	"os"
)

type Data struct {
}

func (self Data) String() string {
	return "data"
}

func main() {
	fmt.Printf("hello %s\n", "world")
	fmt.Println("hello world")

	str := fmt.Sprintf("float %f", 3.14159)
	fmt.Print(str)

	fmt.Fprintln(os.Stdout, "写入指定的输入流")
	//os.Stdout屏幕输出

	fmt.Println("%v\n", Data{}) //String 接口应用

}

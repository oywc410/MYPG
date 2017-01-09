package main

import "fmt"  //我们需要使用fmt包中的Println() 函数

func main() {
	fmt.Println("Hello, world. 你好, 世界! ");
}

/**
	命令编译解释:
		gp run hello.go #直接运行(编译+连接+运行)
	
		go build hello.go #产生编译文件
		./hello #运行 win为hello
		
		
*/
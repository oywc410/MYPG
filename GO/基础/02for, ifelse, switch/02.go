package main

import "fmt"

func main() {
	v_int := 9

	if v_int < 10 {
		fmt.Println("true")
	}

	// 可以省略条件判断的小括号
	if v_int < 10 {
		fmt.Println("true")
	}

	// 可以定义代码块作用域的局部变量
	if temp := 9; v_int > 0 {
		fmt.Printf("%d / %d = %d", temp, v_int, temp/v_int)
	}

	// 在进行条件判断前，可以执行一些初始化的操作，比如打印一些信息
	// else 必须跟在上一个代码块花括号右半部分后面，否则会有编译错误
	if fmt.Println("Checking number:"); v_int%2 == 0 {
		fmt.Printf("%d is even number\n", v_int)
	} else {
		fmt.Printf("%d is odd number\n", v_int)
	}
}
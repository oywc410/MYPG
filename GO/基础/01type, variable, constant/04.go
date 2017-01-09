package main

import (
	"fmt"
	"unsafe"
)

// 使用 const 关键字同时声明多个常量
const int_x, int_y int = 1, 2

// 常量也可以省略类型名让编译器进行类型推断
const str = "Golang!"

// 通过常量数组声明多个常量，如果不同的常量有相同的值，可以省略类型名和初始值
const (
	a, b = 1001001, 73
	c    = true
	d
)

func main() {
	fmt.Println(int_x, int_y)
	fmt.Println(str)
	fmt.Println(a, b, c, d)

	// 常量的值可以是编译期可以确定值的函数比如 len，unsafe.Sizeof 等的返回值
	const sizeofInt = unsafe.Sizeof(int_x)
	fmt.Printf("sizeofInt: %d", sizeofInt)

	// 和变量不同，未使用的局部常量不会导致编译错误
	const no_use = false

}

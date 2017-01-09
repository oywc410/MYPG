package main

import "fmt"

var (
	g_str  string = "TBB"
	g_Int  int    = 73
	g_Bool bool   = true
)

func main() {

	// 使用 var 关键字声明一个变量，变量类型跟在变量名后面
	var a string = "Golang"
	fmt.Println(a)

	// 声明变量时如果赋予了初始值，则可以省略类型，编译器会自动根据给定的初始化值进行推断
	var d = true
	fmt.Println(d)

	// 使用 var 也可以同时声明多个变量，Golang 会对变量从左至右依次赋值
	// 多个变量属于不同类型时
	var v_int, int_value = "v_int is: ", 2
	fmt.Println(v_int, int_value)
	// 多个变量属于相同类型时
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 如果声明变量的时候没有给定初始化值，变量会被初始化为其类型的空值
	var e int
	fmt.Println(e)

	// 在函数内部声明局部变量时，:= 语法可以用来简化已有初始化值的变量声明，下面的语句和 var f string = "short" 是等价的
	// 使用这种用法的时候需要很小心，如果不小心忘掉了 ：符号，就有可能变成对同名的全局变量赋值了。
	f := "short term"
	fmt.Println(f)

	// 声明一个空的 slice 变量并打印
	v_slice := make([]string, 3)
	fmt.Println("Empty slice: ", v_slice)
	v_slice = nil
	fmt.Println("nil slice: ", v_slice)

	fmt.Println(g_Int, g_Bool, g_str)
}

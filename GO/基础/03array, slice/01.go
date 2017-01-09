package main

import "fmt"

// 值传递方式传递数组对象
func arrayTestByValue(arr [1]int) {
	fmt.Println("Passed by value, Address of arr[0] is: ", &arr[0])
}

// 引用传递方式传递数组对象
func arrayTestByref(arr *[1]int) {
	fmt.Println("Passed by ref, Address of arr[0] is: ", &arr[0])
}

func main() {

	// 定义空数组，方括号中的 ... 不能省略，否则就是 slice 对象了。
	v_IntArray := [...]int{}
	fmt.Printf("%T\n", v_IntArray)
	fmt.Println("len(v_IntArray): ", len(v_IntArray))
	fmt.Println("cap(v_IntArray): ", cap(v_IntArray))

	// 声明定长数组，如果没有给定初始值，默认初始化元素为零值，对于 int 类型，就是 0
	v_IntArrayOf5 := [5]int{}
	fmt.Println("len(v_IntArrayOf5: ", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)
	v_IntArrayOf5[2] = 3
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)

	// 声明数组的同时进行初始化
	v_IntArrayOf5 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("len(v_IntArrayOf5: ", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)

	// 可以给数组中的部分元素给定初始值，下面的方式中数组前两个元素会被初始化为指定值
	v_IntArrayOf5 = [5]int{1, 2}
	fmt.Println("len(v_IntArrayOf5: ", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)

	// 可以使用索引进行个别元素的初始化
	v_IntArrayOf5 = [5]int{1: 2, 3: 4}
	fmt.Println("len(v_IntArrayOf5: ", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)

	// 可以进行分段初始化
	v_IntArrayOf5 = [5]int{0: 1, 2, 3: 4, 5}
	fmt.Println("len(v_IntArrayOf5: ", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)

	// 数组进行参数传递的时候是值拷贝方式
	// arrayTestByValue 函数中的数组在内存中的地址和原始地址不相同了。
	// arrayTestByref 函数中数组在内存中的地址和原始数组相同
	v_IntArrayOf1 := [1]int{}
	fmt.Println("Original Address of v_IntArrayOf1[0] is: ", &v_IntArrayOf1[0])
	arrayTestByValue(v_IntArrayOf1)
	arrayTestByref(&v_IntArrayOf1)

	// Golang 支持多维数组，下面是 2 维数组的例子
	v_IntArrayOf23 := [2][3]int{}
	fmt.Println("v_IntArrayOf23: ", v_IntArrayOf23)

}

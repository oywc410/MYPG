package main

import "fmt"

func main() {

	// 通过初始化表达式明确定义 slice 里面每个元素
	v_slice := []int{0, 0, 0}
	fmt.Println("v_slice: ", v_slice)

	// 修改 slice 中的元素值
	v_slice[1] = 10
	fmt.Println("v_slice: ", v_slice)

	// 通过内置函数 make 初始化 slice。最后一个参数表示 Capacity，如果省略，则默认值等于第二个参数 length。
	// 编译时编译器会将 make 翻译为 makeslice 进行对象创建并返回对象。
	v_slice = make([]int, 3, 3)
	v_slice[2] = 10
	fmt.Println("v_slice: ", v_slice)

	// 通过内置函数 new 初始化 slice，省略了 Capacity。编译器创建对象后返回对象指针
	// 下面的官方文档中说这种用法基本没什么用
	// https://golang.org/doc/effective_go.html#slices
	vp_slice := new([]int)
	fmt.Println("vp_slice: ", vp_slice)
	fmt.Println("vp_slice len: ", len(*vp_slice))
	fmt.Println("vp_slice capacity: ", cap(*vp_slice))
}

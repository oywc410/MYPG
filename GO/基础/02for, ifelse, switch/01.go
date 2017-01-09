package main

import "fmt"

func main() {
	v_str := "post hoc"
	fmt.Println("Test str: ", v_str)

	// 类似于 C/C++ 语言的传统使用方式
	for i := 0; i < 4; i++ {
		fmt.Printf("v_str[%d]: %q\t", i, v_str[i])
	}
	fmt.Println()

	// 同时对多个局部变量进行赋值
	for i, length := 0, len(v_str); i < length; i++ {
		fmt.Printf("v_str[%d]: %q\t", i, v_str[i])
	}
	fmt.Println()

	// 配合 range 使用
	// 下面代码中 elem 是 v_str 中每个元素值的副本，在循环体内可以对其进行任意修改，不会影响到原字符串
	for i, elem := range v_str {
		elem = elem - 32
		fmt.Printf("v_str[%d]: %q\t", i, elem)
	}
	fmt.Println("\nTest str: ", v_str)

	// 只保留条件判断：
	fmt.Println("Beginning 4 chars in v_str: ")
	j := 0
	for j < 4 {
		fmt.Printf("v_str[%d]: %q\t", j, v_str[j])
		j++
	}
	fmt.Println()

	// 无条件判断限制的循环，需要在合适的时候从内部 break 或者 return，否则无限循环
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Printf("count: %d\n", count+1)
		count++
	}

	v_IntArray := [...]int{1, 2, 3, 4, 5}

	// 注意：如果 range 处理的是一个指针的数组，那么在循环体内通过指针修改其指向的对象是可以的。
	// range 表达式处理数组，代码块中的改变不会影响到原始数组
	for i, elem := range v_IntArray {
		elem *= 100
		fmt.Println(i, elem)
	}
	fmt.Println("After range []T", v_IntArray)

	// range 表达式处理数组的指针，代码块中的改变也不会影响到原始数组
	for i, elem := range &v_IntArray {
		elem *= 100
		fmt.Println(i, elem)
	}
	fmt.Println("After range *[]T", v_IntArray)

	// range 表达式处理指针的数组，代码块中对指针指向内容的改变会改变原数组中指针指向的内容
	v_Pointer2IntArray := [...]*int{&v_IntArray[0], &v_IntArray[1], &v_IntArray[2], &v_IntArray[3], &v_IntArray[4]}
	for i, elem := range v_Pointer2IntArray {
		*elem *= 100
		fmt.Println(i, *elem)
	}
	fmt.Println("After range []*T, the values being pointed by pointers in v_Pointer2IntArray is:")
	for _, elem := range v_Pointer2IntArray {
		fmt.Printf("%d\t", *elem)
	}
}

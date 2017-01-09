package main

import "fmt"

func main() {
	// 可以替代 if-else 语句
	v_int := 73
	switch {
	case v_int < 100:
		fmt.Println(v_int, "is less than 100")
	case v_int >= 100:
		fmt.Println(v_int, "is greater than or enqual to 100")
	}

	// 分支可以使用变量
	v_intArr := []int{1, 2, 3}
	v_int = 2
	switch v_int {
	case v_intArr[1]:
		println("hello")
	case len(v_intArr):
		println("b")
	default:
		println("c")
	}

	// 使用 fallthrough 之后从执行下一个分支，而且不判断条件。注意 20 是偶数
	v_int = 20
	switch v_int {
	case 20:
		println(v_int, "is even, but:")
		fallthrough
	case 0:
		fmt.Println(v_int, "is odd number")
	}
}

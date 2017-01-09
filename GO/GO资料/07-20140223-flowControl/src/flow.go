package main

import (
	"fmt"
)

func main() {

	c := 1;

	//if语句
	if c < 5 {
	}
	
	//选择语句
	switch c {
		case 0:
			fmt.Printf("0")
		case 2:
			fallthrough//调至3(即输出3)
		case 3:
			fmt.Printf("3");
		case 4, 5, 6:
			fmt.Printf("4, 5, 6");
		default://可与不写
			fmt.Printf("Default");
	}
	
	//循环语句
	sum := 0
	for i := 0; i < 10; i++ {
		sum +=i
	}
	
	sum = 0
	
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	
	//在表达式中进行多重赋值
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("test");
	
	
	
	//指定中断循环的目标
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop //指定中断目标
			}
		}
	}
	JLoop:
	
	
	
}
//跳转语句
func myfunc() {
	i := 0
		HERE:
		fmt.Println(i)
		i++
	if i < 10 {
		goto HERE
	}
}
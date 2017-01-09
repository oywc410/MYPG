package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 5, 10)
	
	fmt.Println("len(mySlice):", len(mySlice))//返回数组个数  5
	fmt.Println("cap(mySlice):", cap(mySlice))//返回数组切片分配的空间 10
	
	//为数组切片新增元素
	mySlice = append(mySlice, 1, 2, 3)
	
	//合并数组切片
	mySlice2 := make([]int, 10, 12)
	mySlice2 = append(mySlice2, mySlice...)
	fmt.Println(mySlice2);
	
	//内容复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
}
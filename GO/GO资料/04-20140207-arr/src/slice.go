package main

import (
	"fmt"
)

/**
数组切片实例
*/
func main() {
	
	//先预定一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	//基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	//mySlice = myArray[:]	基于myArray所有元素创建数组切片
	//mySlice = myArray[5:]	基于从第五个元素开始的所有元素创建数组切片
	
	
	fmt.Println("Elements of myArray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	
	fmt.Println("\nElements of mySlice:")
	
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	
	fmt.Println()

	//直接创建数组切片
	//创建元素为5个初始值为0的数组切片
	mySlice1 := make([]int, 5)
	//创建元素为5个初始值为10的数组切片
	mySlice2 := make([]int, 5, 10)
	//直接创造并初始化包含5个元素的数组切片
	mySlice3 := []int{1, 2, 3, 4, 5}
	
	
	//元素的遍历
	for i := 0; i<len(mySlice); i++ {
		fmt.Println("mySlice[", i, "] =", mySlice[i])
	}
	
	for i, v := range mySlice {
		fmt.Println("mySlice[", i, "] =", v)
	}
	
	fmt.Println(mySlice1, mySlice2, mySlice3);
	
}
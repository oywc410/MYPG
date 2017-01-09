package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func main() {
	p1 := new(Person)
	p2 := make([]Person, 10)

	//指针类型
	fmt.Println(p1)

	//make对象(切片)
	fmt.Println(p2)
}

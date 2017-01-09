package main

import (
	"fmt"
)

type Data int

func (data Data) String() string {
	return string(data)
}

func main() {
	checkInterface()
	setFunc()
}

func checkInterface() {
	//让编译器检查,以确保某个类型实现接口
	var _ fmt.Stringer = (*Data)(nil)
}

func setFunc() {
	//让函数直接实现接口
	var t Tester = FuncDo(func() { fmt.Println("test") })
	t.Do()
}

type Tester interface {
	Do()
}

type FuncDo func()

func (self FuncDo) Do() { self() }

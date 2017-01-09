package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}

	return fmt.Sprintf(s, c)
}

//动态调用方法
//非导出方法无法调用,甚至无法透过指针操作,因为接口类型中没有该方法地址
func main() {
	func1()
	func2()
}

func func1() {
	d := new(Data)
	v := reflect.ValueOf(d)

	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		//调用函数
		out := m.Call(in)

		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	fmt.Println("--------------------")

	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}

func func2() {
	d := new(Data)
	v := reflect.ValueOf(d)

	m := v.MethodByName("Sum")

	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}), //打包成slice
	}

	out := m.CallSlice(in)

	for _, v := range out {
		fmt.Println(v.Interface())
	}
}

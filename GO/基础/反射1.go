package main

import (
	"fmt"
	"reflect"
)

//http://golangtc.com/t/53317f90320b5261e0000058
//http://golanghome.com/post/105

type T struct {
	A int
	B string
}

func (t *T) aaa() {
	fmt.Println(t)
}

func main() {
	structFunc()
	typeFunc()
	setFunc2()
}

func structFunc() {
	//获取对象中的值
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v, isSet %s\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface(), f.CanSet())
	}
}

func typeFunc() {
	//获取变量值
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var y uint8 = 'x'
	v = reflect.ValueOf(y)
	fmt.Println("type:", v.Type())                            // uint8
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true
	y = uint8(v.Uint())
	// v.Uint 返回一个 uint64
}

func setFunc() {
	//设置属性
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//v.SetFloat(7.1) // Error: will panic.
	fmt.Println("settability of v:", v.CanSet()) //判断是否可以被设置
}

func setFunc2() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // 注意：获取 X 的地址。
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	//v可设置
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.7)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

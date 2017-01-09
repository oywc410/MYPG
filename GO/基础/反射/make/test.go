package main

import (
	"fmt"
	"reflect"
)

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func Make(T reflect.Type, fptr interface{}) {
	//实际创建slice的包装函数
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.MakeSlice(
			reflect.SliceOf(T),
			int(in[0].Int()), 
			int(in[1].Int())
	}

	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), swap)

	fn.Set(v)
}

func main() {

}

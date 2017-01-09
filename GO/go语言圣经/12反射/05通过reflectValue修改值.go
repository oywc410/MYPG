package main

import (
	"reflect"
	"fmt"
	"os"
)

func main() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)

	//判断是否可以取得地址
	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true"

	//修改值方法1
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"

	//修改值方法2
	d.Set(reflect.ValueOf(4))
	d.SetInt(12)
	fmt.Println(x) // "12"

	//修改时类型限制
	x = 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)                     // OK, x = 2
	rx.Set(reflect.ValueOf(3))       // OK, x = 3
	rx.SetString("hello")            // panic: string is not assignable to int
	rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	ry.SetInt(2)                     // panic: SetInt called on interface Value
	ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
	ry.SetString("hello")            // panic: SetString called on interface Value
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"

	//访问私有变量
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.Int()) // "1"
	fd.SetInt(2)          // panic: unexported field

	fmt.Println(fd.CanAddr(), fd.CanSet()) // "true false"
	//fd.CanAddr() 判断地址是否可被访问
	//fd.CanSet() 判断地址可被修改
}
package main

import (
	"fmt"
	"text/template"
	"unsafe"
)

func main() {
	func1()
	func2()
}

/*
//不能对指针进行直接运算
func func1() {
	u := uint32(32)
	i := int32(1)
	fmt.Println(&u, &i) //打印地址
	p := &i             //p的类型是*int32
	p = &u              //&u的类型是*uin32,于p的类型不同,不能赋值
	p = (*int32)(&u)    //这种类型转换语也是无效的
	fmt.Println(p)
}
*/

func func1() {
	u := uint32(32)
	i := int32(1)
	fmt.Println(&u, &i)
	p := &i
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(p, *p)
}

type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}

//突破私有变量
func func2() {
	t := template.New("foo")
	fmt.Println(&t)
	//Pointer类型用于表示任意类型的指针 16进制
	fmt.Println(unsafe.Pointer(t))
	//uintptr是整型,可以足够保存指针的值得范围,在32平台下为4字节,在64位平台下是8字节
	fmt.Println(uintptr(unsafe.Pointer(t)))
	p := (*MyTemplate)(unsafe.Pointer(t))
	p.name = "Bar" //突破私有变量
	fmt.Println(p, t)
}

package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main() {
	func1()
	func2()
}

//类型判断
func func1() {
	var o interface{} = &User{1, "Tom"}

	//可判断接⼝口对象是否某个具体的接⼝口或类型
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}
	u := o.(*User)
	fmt.Println(u)
}

//类型判断
func func2() {
	var o interface{} = &User{1, "Tom"}
	//此处的switch不支持fallthroug
	switch v := o.(type) {
	case nil: // o == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface
		fmt.Println(v)
	case func() string: // func
		fmt.Println(v())
	case *User: // *struct
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}

type Stringer interface {
	String() string
}

type Printer interface {
	String() string
	Print()
}

func (self *User) Print() {
	fmt.Println(self.String())
}

//接口转换
func func3() {
	//超集接口对象可转换为⼦子集接⼝口，反之出错。
	var o Printer = &User{1, "Tom"}
	var s Stringer = o
	fmt.Println(s.String())
}

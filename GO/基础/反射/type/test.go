package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}

type Admin struct {
	User
	title string
}

func main() {
	fmt.Println("------------func1----------")
	func1()
	fmt.Println("------------func2----------")
	func2()
	fmt.Println("------------func3----------")
	func3()
	fmt.Println("------------func4----------")
	func4()
	fmt.Println("------------func5----------")
	func5()
	fmt.Println("------------func6----------")
	func6()
	fmt.Println("------------func7----------")
	func7()
	fmt.Println("------------func8----------")
	func8()
}

//普通对象
func func1() {
	var u Admin

	t := reflect.TypeOf(u)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

//指针
func func2() {
	u := new(Admin)
	t := reflect.TypeOf(u)

	if t.Kind() == reflect.Ptr {
		//如果是指针，应该先使⽤用 Elem ⽅方法获取目标类型，指针本⾝身是没有字段成员的
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

//方法
func (*User) ToString() {}

func (Admin) test() {}

func func3() {
	var u Admin
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("----value interface----")
	methods(reflect.TypeOf(u))

	fmt.Println("----pointer interface----")
	methods(reflect.TypeOf(&u))
}

//可直接用名称或区号访问字段,包括多级序号访问嵌入文字成员 (继承)
func func4() {
	var u Admin
	t := reflect.TypeOf(u)

	f, _ := t.FieldByName("title")
	fmt.Println(f.Name)

	f, _ = t.FieldByName("User")
	fmt.Println(f.Name)

	f, _ = t.FieldByName("Username") //直接访问嵌入字段成员
	fmt.Println(f.Name)

	f = t.FieldByIndex([]int{0, 1}) //Admin[0]->User[1]->age
	fmt.Println(f.Name)
}

//ORM应用
type UserOrm struct {
	Name string `field:"username" type:"navrchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
}

func func5() {
	var u UserOrm

	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("Name")

	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
}

//复合类型
var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func func6() {
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)
	//chan<-string

	m := reflect.MapOf(String, Int)
	fmt.Println(m)
	//map[string]int

	s := reflect.SliceOf(Int)
	fmt.Println(s)
	//[]int

	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)
	//*struct{Name string}

	f := reflect.TypeOf(make(chan int)).Elem()
	//⽅方法 Elem 可返回复合类型的基类型。
	fmt.Println(f)
	//int
}

//判断是否实现了某个接口
type Data struct {
	b byte
	x int32
}

func (*Data) String() string {
	return ""
}

func func7() {
	var d *Data
	t := reflect.TypeOf(d)

	//没法直接获取接口类型,创建一个空指针对象
	//这样传递TypeOf转换成interface{}
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	fmt.Println(t.Implements(it))
}

//获取对其信息对于内存自动分析是有用的
func func8() {
	var d Data

	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align()) //sizeof,以及最宽字段的対齐摸数
	//8 4

	f, _ := t.FieldByName("b")
	fmt.Println(f.Type.FieldAlign()) //字段対齐
	//1

	f, _ = t.FieldByName("x")
	fmt.Println(f.Type.FieldAlign())
	//4
}

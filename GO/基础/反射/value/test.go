package main

import (
	"fmt"
	"reflect"
	"unsafe"
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
	fmt.Println("-----func1--------")
	func1()
	fmt.Println("-----func2--------")
	func2()
	fmt.Println("-----func3--------")
	func3()
	fmt.Println("-----func4--------")
	func4()
	fmt.Println("-----func5--------")
	func5()
}

//struct
func func1() {
	u := &Admin{User{"Jak", 23}, "NT"}
	v := reflect.ValueOf(u).Elem()

	fmt.Println(v.FieldByName("title").String())
	fmt.Println(v.FieldByName("age").Int())
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int())

	//非导出类型判断
	f := v.FieldByName("age")
	if f.CanInterface() {
		fmt.Println(f.Interface())
	} else {
		fmt.Println(f.Int())
	}

	//不在值
	u1 := User{}
	v1 := reflect.ValueOf(u1)
	f1 := v1.FieldByName("a")
	fmt.Println(f1.Kind(), f1.IsValid())

	//空值
	var p *int
	var x interface{} = p
	fmt.Println(x == nil)
	v = reflect.ValueOf(p)
	fmt.Println(v.Kind(), v.IsNil())
}

//slice,array,map
func func2() {
	v := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, v.Len(); i < n; i++ {
		fmt.Println(v.Index(i).Int())
	}

	fmt.Println("-------------")

	v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int())
	}
}

//修改字段
func func3() {
	u := User{"Jack", 23}

	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	//判断是否可以修改值
	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
}

//强行修改私有字段(修改指针)
func func4() {
	u := User{"Jack", 23}
	p := reflect.ValueOf(&u).Elem()

	//共有变量
	p.FieldByName("Username").SetString("Tom")

	//私有变量
	f := p.FieldByName("age")

	//判断是否可以获取地址
	if f.CanAddr() {
		//获取指针地址
		fmt.Println(f.UnsafeAddr())                   //10进制
		fmt.Println(unsafe.Pointer(f.UnsafeAddr()))   //16
		age := (*int)(unsafe.Pointer(f.UnsafeAddr())) //创建指针(地址为age,即可直接修改地址)
		//age := (*int)(unsafe.Pointer(f.Addr().Pointer())) 与上面结果相同
		*age = 88
	}

	// 注意 p 是 Value 类型，需要还原成接⼝口才能转型。
	fmt.Println(u, p.Interface().(User))
}

//复合类型修改示例
func func5() {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)

	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

	fmt.Println(v2.Interface())
	//fmt.Println(v2.Interface().([]int))

	fmt.Println("-----------------")

	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()

	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) //update
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) //add

	fmt.Println(v.Interface().(map[string]int))
}

package main

import (
	"fmt"
)

//声明借口
type Animal interface {
	Grow()
	Move(string) string
	aaa()
}

//声明对象
type Cat struct {
	Name     string
	Age      int32
	Location string
}

//实现接口方法
func (cat *Cat) Grow() {
	cat.Age++
}

//实现接口方法  * 由如php & 的用法
func (cat *Cat) Move(newLo string) string {
	fmt.Printf(newLo)
	oldAddr := cat.Location
	cat.Location = newLo
	return oldAddr
}

func (cat *Cat) aaa() {

}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	//将对象放入接口
	animal, ok := interface{}(&myCat).(Animal)
	//调用接口方法
	animal.Grow()
	animal.Move("test")
	fmt.Printf("%v, %v\n", ok, animal)
}

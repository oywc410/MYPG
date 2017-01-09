package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer //嵌入接口
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func Print(t Printer) {
	t.Print()
}

func main() {
	var t Printer = &User{1, "Tom"}
	//User必须实现Print, String
	Print(t)
}

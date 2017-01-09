package main

import (
	"fmt"
)

type Tester struct {
	s interface {
		String() string
	}
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func main() {
	t := Tester{&User{1, "Tom"}}
	//User必须实现s接口
	fmt.Println(t.s.String())
}

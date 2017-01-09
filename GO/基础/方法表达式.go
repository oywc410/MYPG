package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (self User) Test() {
	fmt.Printf("%p, %v\n", &self, self)
}

func main() {
	u := User{1, "Tom"}

	mValue := u.Test //立即复制 receiver，因为不是指针类型，不受后续修改影响。

	mExpression := (*User).Test

	u.name = "aaa"

	u.Test()
	mValue()        // 隐式传递 receiver
	mExpression(&u) // 显式传递 receiver
	u.Test()
	/**
	0xc082002720, {1 aaa}
	0xc0820027a0, {1 Tom}
	0xc0820027e0, {1 aaa}
	0xc082002820, {1 aaa}
	*/
}

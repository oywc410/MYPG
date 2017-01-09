package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func main() {
	u := User{1, "Tom"}
	var i interface{} = u

	u.id = 2
	u.name = "Jack"

	//i.(User).id = 7 error

	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User))

	var j interface{} = &u

	j.(*User).name = "1111"

	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)

	fmt.Println(u)

	/**
	{2 Jack}
	{1 Tom}
	{1 Tom}
	&{2 1111}
	{2 1111}
	*/
}

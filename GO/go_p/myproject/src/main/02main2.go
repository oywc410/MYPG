package main

import (
	"fmt"
)

type Task struct {
	ID     int
	Detail string
	done   bool
	*User
}

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func NewTask(id int, detail, firstName, lastName string) (task *Task) {
	task = &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}

	return
}

func main() {
	//对象嵌套
	task := NewTask(1, "aaa", "bbbb", "ccc")
	fmt.Println(task.FirstName)

	//类型转换
	var i uint8 = 3
	var j uint32 = uint32(i)
	fmt.Println(j)

	var s string = "abc"
	var b []byte = []byte(s)
	fmt.Println(b)

}

//抽象类嵌套
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

//类型断言

func Print(value interface{}) {
	s, ok := value.(string)
	if ok {
		fmt.Println("value is string: ", s)
	} else {
		fmt.Println("value is not string", s)
	}
}

//类型判断
func Print(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %d\n", v)
	case Stringer:
		fmt.Printf("value is Stringer: %s\n", v)
	}
}

package main

import (
	"errors"
	f "fmt"
	"gosample"
	"log"
	"os"
)

var foo, bar, buz string = "foo", "bar", "buz"

var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
)

func main() {
	f.Println(gosample.Message)
	f.Println(foo, bar, buz)
	f.Println(a, b, c)

	message := "hello world"
	f.Println(message)

	n := 0
	for {
		n++
		if n > 10 {
			break
		}
		if n%2 == 0 {
			continue
		}
		f.Println(n)
	}

	n = 10

	switch n {
	case 15:
		f.Println("FizzUbzz")
	case 5, 10:
		f.Println("nizz")
	case 3, 6, 9:
		f.Println("Fizz")
	default:
		f.Println(n)
	}

	b, err := div(10, 1)
	if err != nil {
		log.Fatal(err)
	}
	f.Println(b)

	var arr [4]string
	arr[0] = "a"
	arr[1] = "b"

	f.Println(arr)
	//数组
	arr1 := [4]string{"a", "b", "c", "d"}
	arr2 := [...]string{"a", "b", "c", "d", "c"}
	f.Println(arr1)
	f.Println(arr2)

	fn(arr1)
	//错误
	//fn(arr2)

	var s []string
	//为数组添加新值
	s = append(s, "a", "b", "c")
	f.Println(s)

	s1 := []string{"a", "b"}
	s2 := []string{"c", "d"}
	//合并数组
	s1 = append(s1, s2...)
	f.Println(s1)
	//遍历数组
	for i, t := range s1 {
		f.Println(i, t)
	}
	//切片
	f.Println(arr2[2:4])
	f.Println(arr2[:])
	f.Println(arr2[2:])
	//map
	var month map[int]string = map[int]string{}

	month[1] = "J"
	month[2] = "a"

	month = map[int]string{
		1: "j",
		2: "F",
	}

	f.Println(month)

	delete(month, 1)
	//map中元素存在判断
	mat, m_ok := month[1]
	if m_ok {
		f.Println(mat)
	} else {
		f.Println(m_ok)
	}

	var i2 int = 10
	callByValue(i2)
	f.Println(i2)
	callByRef(&i2)
	f.Println(i2)

	file, err := os.Open("hello.go")
	if err != nil {
		f.Println(err)
	}
	//不管是否发生错误 一定执行defer
	defer file.Close()

	defer div(7, 5)

	defer func() {
		//捕获错误
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var task Task = Task{
		ID:     1,
		Detail: "aaaa",
		done:   true,
	}
	f.Println(task.ID)

	Finish(&task)

	f.Println(task.ID)

	task2 := &Task{
		ID:     1,
		Detail: "aaaa",
		done:   true,
	}

	task2.Finish2()

	f.Println(task2.ID)

	f.Println(task, task2)

	var d string = "aaaa"
	Print(d)

	q := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		if i >= len(q) {
			//定义致命性错误
			//panic(errors.New("index out of range"))
		}
	}
}

func Print(value interface{}) {
	switch v := value.(type) {
	case string:
		f.Printf("value is string: %s\n", v)
	case int:
		f.Printf("value is int: %d\n", v)
	case Stringer:
		f.Printf("value is Stringer: %s\n", v)
	}
}

type Task struct {
	ID     int
	Detail string
	done   bool
}

func Finish(task *Task) {
	task.ID = 2
}

func (task *Task) Finish2() {
	task.ID = 3
}

type Stringer interface {
	String() string
}

//任意变量
func Do(e interface{}) {

}

/*
//任意长度变量
func PrintlnMy(a ...interface{}) (n int, err error) {

}
*/

func div(i, j int) (int, error) {
	if j == 0 {
		//自定义错误
		return 0, errors.New("divide by zero")
	}
	return i / j, nil
}

func fn(arr [4]string) {
	f.Println(arr)
}

func fn2(arr ...string) {
	f.Println(arr)
}

func callByValue(i int) {
	i = 20
}

func callByRef(i *int) {
	*i = 20
}

package funcs

import (
	"errors"
	"fmt"
)

//函数声明
func Add(a int, b int) (ret int, err error) {
	if a < 0 || b  < 0 {
		err = errors.New("Error");
		return
	}
	
	return a + b, nil	//支持多重返回值
}

//不定参数类型(传入参数个数不定)
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg);
	}
}


//不定参数的传递
func myfunc2(args ...int) {
	//按原样传递
	myfunc3(args...)
	//传递切片
	myfunc3(args[1:]...)
}

func myfunc3(args ...int) {
}

//任意类型的不定参数
func Printf(format string, args ...interface{}) {
}

//函数的多返回值
func test() {
	a, _ := testFunc()
	//a = 1
}

func testFunc() (a int, b int) {

	//匿名函数
	f := func(x, y int) int {
		return x + y
	}
	
	//直接执行
	func(ch chan int) {
		ch <- ACK
	} (reply_chan) // 花括号后直接跟参数列表表示函数调用

	return 1, 2;
}


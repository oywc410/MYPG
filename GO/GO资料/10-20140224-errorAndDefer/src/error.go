package src

import (

)

/*
	func Foo(param int) (n int, err error) {
	}
	
	n, err := Foo(0)
	
	if err != nil {
		//错误处理
	} else {
		//使用返回值n
	}
*/


//自定义error类型
type PathError struct {
	Op string
	Path string
	Err error
}

//声明PathError 为 error来传递
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}


func Stat(name string) (fi FileInfo, err error) {
	var stat syscall.Stat_t
	err = syscall.Stat(name, &stat)
	
	if err != nil {
		return nil, &PathError{"stat", name, err}
	}
	
	return fileInFromStat(&stat, name), nil
}

/**
fi, err := os.Stat("a.txt")
if err != nil {
	if e, ok := err.(*os.PathError); ok && e.Err != nil {
		// 获取PathError类型变量e中的其他信息并处理
	}
}

*/

/**
	p64  ????

	Go语言引入了两个内置函数panic()和recover()以报告和处理运行时错误和程序中的错误场景：
	func panic(interface{})
	func recover() interface{}
	
	当在一个函数执行过程中调用panic()函数时，正常的函数执行流程将立即终止，但函数中
	之前使用defer关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，并导致
	逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止。错误信息将被报
	告，包括在调用panic()函数时传入的参数，这个过程称为错误处理流程。
	
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()
	foo()
*/
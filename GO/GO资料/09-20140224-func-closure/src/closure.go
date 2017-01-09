package main

import (
	"fmt"
)

//闭包实例
func main() {
	var j int = 5
	
	a := func()(func()) {
		/*
			在上面的例子中，变量a指向的闭包函数引用了局部变量i和j，i的值被隔离，在闭包外不
			能被修改，改变j的值以后，再次调用a，发现结果是修改过的值。
		*/
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	
	a()
	
	j *= 2
	a();
}
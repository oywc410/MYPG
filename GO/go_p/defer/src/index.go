package main

import (
	"fmt"
)

func main() {
	deferIt2()
	fmt.Println("\n")
	deferIt3()
	fmt.Println("\n")
	deferIt4()
}


func deferIt2() {
    for i := 1; i < 5; i++ {
        defer fmt.Print(i)
    }
}

func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)//正常输出
        return i * 10//延迟输出 (函数结束时输出)
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}

func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func() {
            fmt.Print(i)//只输出为5
        }()
    }
}    

 
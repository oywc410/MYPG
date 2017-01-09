package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)//增加  (开启阻塞)
    go func() {
        fmt.Println("Go!")
        wg.Done()//减少			2
    }()
    go func() {
        fmt.Println("Go!")
        wg.Done()//				1
    }()
    go func() {
        fmt.Println("Go!")
        wg.Done()//				0
    }()
    wg.Wait()//阻塞结束  为0时结束阻塞
}
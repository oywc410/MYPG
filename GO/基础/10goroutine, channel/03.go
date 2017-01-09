package main

import (
	"sync"
	"fmt"
	"time"
)

func ExamplaCond() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	//用作向goroutine 内发送等待与结束命令
	c := sync.NewCond(&mu)

	for i:=0; i<10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer mu.Unlock()

			fmt.Printf("waiting %d\n", i)
			mu.Lock()
			c.Wait()//等待
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i:=0; i<10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("signaling!\n")

		c.Signal()//发送继续运行信息
	}

	wg.Wait()
}

func main() {
	ExamplaCond()
}
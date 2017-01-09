package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()

	return c
}

func func1() {
	//用简单⼯工厂模式打包并发任务和 channel
	t := NewTest()
	println(<-t)
}

func func2() {
	//用 channel 实现信号量 (semaphore)。
	wg := sync.WaitGroup{}
	wg.Add(3)

	sem := make(chan int, 1)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()

			sem <- 1

			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}

			<-sem
		}(i)
	}

	wg.Wait()
}

func main() {
	func1()
	func2()
}

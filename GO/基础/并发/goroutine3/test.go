package main

import (
	"runtime"
	"sync"
)

type aaa int

func main() {
	var wg sync.WaitGroup
	//wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				//，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等待下次被调度执⾏行
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		println("hello, world!")
	}()

	wg.Wait()
}

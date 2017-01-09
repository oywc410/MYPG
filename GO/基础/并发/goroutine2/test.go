package main

import (
	"runtime"
	"sync"
)

func main() {
	//runtime.Goexit() 终止当前goroutine 并确保defer执行

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer println("A.defer")

		func() {
			defer println("B.defer")
			runtime.Goexit() // 终止当前 goroutine
			println("B")     // 不会执⾏行
		}()

		println("A") // 不会执⾏行
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	w := make(chan bool)
	c := make(chan int, 2)

	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			//AfterAfter会在另一线程经过时间段d后向返回值发送当时的时间
			fmt.Println("timeout.")
		}

		w <- true
	}()

	<-w
}

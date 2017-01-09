package main

import (
	"time"
	"log"
)

func main(){
	var ch1, ch2 chan int

	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()

	//定期実行
	for {
		//select time out
		select {
		case <-ch1:
		case <-ch2:
		case <-ticker.C:
			log.Println("time out") // call some logging function logState
		default: // no value ready to be received
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	data := make(chan int, 3) //缓存区可以储存3个元素
	exit := make(chan bool)

	//在缓存区未满前,不会阻塞
	data <- 1
	data <- 2
	data <- 3

	go func() {
		for d := range data { //在缓存区未空前,不会堵塞
			fmt.Println(d)
		}

		/*
			for {
				//还可用 ok-idiom 模式判断 channel 是否关闭
				if d, ok := <-data; ok {
					fmt.Println(d)
				} else {
					break
				}
			}
		*/

		exit <- true
	}()

	data <- 4
	data <- 5
	close(data)

	<-exit
}

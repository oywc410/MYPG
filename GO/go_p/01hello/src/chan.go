package main

import (
	"fmt"
	"time"
)

//只准输入的通道类型
type Sender chan<- int

//只准输出的通道类型
type Receiver <-chan int

func main() {
	//双向通道类型
	var myChannel = make(chan int, 2) //类型,大小
	fmt.Println(cap(myChannel))
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}

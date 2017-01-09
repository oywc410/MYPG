package main

import (
	"./mytcp"
	"fmt"
	"time"
)

func main() {
	alltest()
}

func alltest() {

	con := 200

	chanTest := make(chan int8, 1000)

	for i := 0; i < con; i++ {

		if i % 100 == 0 {
			time.Sleep(time.Microsecond * 500000)
		}

		go func() {
			mytcp.ClientStart(func(code int8) {
				chanTest <- code
			})
		}()
	}



	fmt.Println("-----------------------------------------")

	var stop1, stop2, stop3, stop4 int

	for j := 0; j < con; j++ {
		switch <-chanTest {
		case 1:
			stop1++
		case 2:
			stop2++
		case 3:
			stop3++
		case 4:
			stop4++
		}
	}

	fmt.Println(stop1, stop2, stop3, stop4)
}
package main

import "errors"

func test() {
	errCh := make(chan error, 1)
	outputCh := make(chan string)
	go func () {
		defer close(errCh)
		defer close(outputCh)
		errCh <- errors.New("error")
	}()

	select {
	case t, cl := <- errCh:
		_ = t
		_ = cl
		if !cl {
			println("errCh", t, cl)
		}
	//println(t, cl)
	case <- outputCh://此处将被输出 (close 也会被输出)
		println("outputCh")
	}
}

func main() {
	for i:=0;i<1000000;i++ {
		test()
	}
}

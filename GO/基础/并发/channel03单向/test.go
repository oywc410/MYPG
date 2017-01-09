package main

func main() {
	c := make(chan int, 3)

	var send chan<- int = c
	var recv <-chan int = c

	send <- 1

	//recv <- 2 error
	<-recv

}

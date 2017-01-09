package main

func main() {
	var test chan int
	<- test //ログ
	//tesr <- 1 //ログ
}

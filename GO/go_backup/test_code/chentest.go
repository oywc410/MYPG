package main

import (
	"fmt"
)

func main() {
	chens := make(chan string, 2)

	chens <- "1"
	chens <- "2"
	chens <- "3"

	fmt.Println(len(chens))
}

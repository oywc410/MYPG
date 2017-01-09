package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	i := 0

	y := 0
	n := 0

	for {
		i++
		line, _, _ := reader.ReadLine()

		if string(line) == "yes" {
			y++
		} else {
			n++
		}

		if i == 5 {
			break
		}
	}

	if y > n {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}

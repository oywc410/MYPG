package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	n, _ := strconv.Atoi(string(line))

	line, _, _ = reader.ReadLine()

	m, _ := strconv.Atoi(string(line))

	t := 0

	for i := 0; i < m; i++ {

		if t/n%2 < 1 {
			fmt.Print("R")
		} else {
			fmt.Print("W")
		}
		t++
	}
}

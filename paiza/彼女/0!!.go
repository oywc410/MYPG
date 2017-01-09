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

	for {

		fmt.Println(n)

		n--

		if n == 0 {
			fmt.Println("0!!")
			break
		}
	}
}

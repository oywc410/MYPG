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

	m := 1

	for i := 1; i <= n; i++ {
		m *= i
	}

	fmt.Println(m)
}

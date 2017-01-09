package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	n, _ := strconv.Atoi(string(line))

	line, _, _ = reader.ReadLine()

	m, _ := strconv.Atoi(string(line))

	fmt.Print(m + n)
}

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

	fmt.Println(100 - n)
}

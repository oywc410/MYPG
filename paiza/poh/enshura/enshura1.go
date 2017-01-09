package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()

	arrStr := string(line)

	for i, value := range arrStr {

		if (i+1)%2 != 0 {
			fmt.Print(string(value))
		}

	}

	fmt.Println("")
}

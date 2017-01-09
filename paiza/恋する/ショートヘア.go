package main

import (
	"bufio"
	"os"
	"io"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	i := 5

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		str := strings.Split(string(line), " ")
		if str[0] != str[1] {
			i--
		}
	}

	if i > 2 {
		fmt.Println("OK")
	} else {
		fmt.Println("NO")
	}
}

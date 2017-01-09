package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	n, _ := strconv.Atoi(string(line))

	a := bytes.Buffer{}

	for i := 0; i < n; i++ {
		//		reader = bufio.NewReader(os.Stdin)

		line, _, _ = reader.ReadLine()

		if i != 0 {
			a.WriteString("_")
		}

		for _, value := range line {
			a.WriteByte(value)
		}

	}

	fmt.Println(a.String())
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var ints int = 0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		constr := string(line)

		if constr == "\\q" {
			break
		}

		bs := strings.Split(constr, " ")

		for _, num := range bs {
			b, err := strconv.Atoi(num)
			if err == nil {
				ints += b
			}
		}
	}

	fmt.Println(ints)
}

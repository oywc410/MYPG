package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var ints []int
	a := false

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		constr := string(line)

		if constr == "\\q" {
			break
		}
		if a == false {
			a = true
		} else {
			b, err := strconv.Atoi(constr)
			if err == nil {
				ints = append(ints, b)
			}
		}
	}
	sort.Ints(ints)

	for _, val := range ints {
		fmt.Println(val)
	}
}

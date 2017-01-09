package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var bs []string
	bbs := []string{}

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		constr := string(line)

		if constr == "\\q" {
			break
		}

		bs = strings.Split(constr, " ")

	}

	con := make(map[string]int)

	for _, str := range bs {
		if _, isEx := con[str]; !isEx {
			bbs = append(bbs, str)
		}
		con[str]++
	}

	for _, str := range bbs {
		fmt.Println(str, con[str])
	}
}

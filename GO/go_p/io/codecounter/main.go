package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var line int

	for {
		//isPreix是否为超长行,超长行时将重复读取
		_, isPreix, err := reader.ReadLine()

		if err != nil {
			break
		}

		if !isPreix {
			line++
		}
	}

	fmt.Println(line)
}

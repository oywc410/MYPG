package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	n, _ := strconv.Atoi(string(line))

	allBook := make([]bool, n, n)

	reader.ReadLine()

	myBook, _, _ := reader.ReadLine()

	var myAllBook, tAllBook map[string]bool

	if len(myBook) > 0 {
		myAllBook = make(map[string]bool)
		myAllBookS := strings.Split(string(myBook), " ")

		for _, value := range myAllBookS {
			myAllBook[value] = true
		}

	}

	reader.ReadLine()

	tBook, _, _ := reader.ReadLine()

	if len(tBook) > 0 {
		tAllBook = make(map[string]bool)
		tAllBookS := strings.Split(string(tBook), " ")
		for _, value := range tAllBookS {
			tAllBook[value] = true
		}
	}

	for key, _ := range allBook {
		i := key + 1
		t := strconv.Itoa(i)

		_, isEx := myAllBook[t]

		if !isEx && tAllBook[t] == true { //存在チェック
			allBook[key] = true
		}
	}

	isGet := false

	for key, value := range allBook {
		if value {
			if isGet {
				fmt.Print(" ")
			}
			isGet = true
			fmt.Print(key + 1)
		}
	}

	if !isGet {
		fmt.Print("None")
	}

}

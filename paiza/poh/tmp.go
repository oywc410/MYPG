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
	topStr := strings.Split(string(line), " ")
	n, _ := strconv.Atoi(topStr[0])
	r, _ := strconv.Atoi(topStr[1])

	R := 2 * r

	for i := 0; i < n; i++ {
		line, _, _ := reader.ReadLine()
		str := strings.Split(string(line), " ")
		h, _ := strconv.Atoi(str[0])
		w, _ := strconv.Atoi(str[1])
		d, _ := strconv.Atoi(str[2])

		if h >= R && w >= R && d >= R {
			fmt.Println(i + 1)
		}
	}

}

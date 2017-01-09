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

	arr1 := strings.Split(string(line), " ")

	line, _, _ = reader.ReadLine()

	arr2 := strings.Split(string(line), " ")

	n1, _ := strconv.ParseFloat(arr1[0], 64)
	m1, _ := strconv.ParseFloat(arr1[1], 64)

	n2, _ := strconv.ParseFloat(arr2[0], 64)
	m2, _ := strconv.ParseFloat(arr2[1], 64)

	if n1/m1 > n2/m2 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}

}

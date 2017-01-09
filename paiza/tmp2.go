package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lineArr []([]int)
var iiArr map[int](map[int]([]int))
var l, n, m int

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	arr := strings.Split(string(line), " ")

	l, _ = strconv.Atoi(arr[0])
	n, _ = strconv.Atoi(arr[1])
	m, _ = strconv.Atoi(arr[2])

	lineArr = make([]([]int), n)
	iiArr = make(map[int](map[int]([]int)))

	for i := 0; i < n; i++ {
		lineArr[i] = make([]int, l)
	}

	for i := 0; i < m; i++ {
		line, _, _ := reader.ReadLine()
		arr := strings.Split(string(line), " ")
		keyL, _ := strconv.Atoi(arr[0])
		keyL--

		keyY, _ := strconv.Atoi(arr[1])
		keyY = l - keyY

		keyY2, _ := strconv.Atoi(arr[2])
		keyY2 = l - keyY2

		keyL2 := keyL + 1

		lineArr[keyL][keyY] = 1
		lineArr[keyL2][keyY2] = 1

		if _, isEx := iiArr[keyL]; !isEx {
			iiArr[keyL] = make(map[int]([]int))
		}

		if _, isEx := iiArr[keyL][keyY]; !isEx {
			iiArr[keyL][keyY] = make([]int, 2, 2)
		}

		iiArr[keyL][keyY][0] = keyL + 1
		iiArr[keyL][keyY][1] = keyY2

		if _, isEx := iiArr[keyL2]; !isEx {
			iiArr[keyL2] = make(map[int]([]int))
		}

		if _, isEx := iiArr[keyL2][keyY2]; !isEx {
			iiArr[keyL2][keyY2] = make([]int, 2, 2)
		}

		iiArr[keyL2][keyY2][0] = keyL
		iiArr[keyL2][keyY2][1] = keyY

	}

	if lineArr[0][0] == 1 {
		fmt.Println(reFunc(iiArr[0][0][0], iiArr[0][0][1]))
	} else {
		fmt.Println(reFunc(0, 0))
	}

}

func reFunc(x, y int) int {

	for i := y + 1; i < l; i++ {

		if lineArr[x][i] == 1 {
			return reFunc(iiArr[x][i][0], iiArr[x][i][1])
		}
	}

	return x + 1
}

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
	nw := 0

	arr1 := make(map[int](map[int]int))

	for i := 0; i < n; i++ {
		arr1[i] = make(map[int]int)
		line, _, _ = reader.ReadLine()

		tmpArr := strings.Split(string(line), " ")
		nw = 0
		for key, value := range tmpArr {
			arr1[i][key], _ = strconv.Atoi(value)
			nw++
		}
	}

	line, _, _ = reader.ReadLine()

	m, _ := strconv.Atoi(string(line))
	mw := 0

	arr2 := make([]int, 0, 100)

	for i := 0; i < m; i++ {
		//arr2[i] = make(map[int]int)
		line, _, _ = reader.ReadLine()

		tmpArr := strings.Split(string(line), " ")
		mw = 0
		for _, value := range tmpArr {
			val, _ := strconv.Atoi(value)
			arr2 = append(arr2, val)
			mw++
		}
	}

	//判断を行う
	for i := 0; i <= n-m; i++ {
		for j := 0; j <= nw-mw; j++ {
			checkArr := make([]int, 0, 100)
			for z := i; z < i+m; z++ {
				for k := j; k < j+mw; k++ {
					checkArr = append(checkArr, arr1[z][k])
				}
			}

			isEx := true
			for key, value := range checkArr {
				if value != arr2[key] {
					isEx = false
					break
				}
			}

			if isEx {
				fmt.Println(i, j)
			}

		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var slip map[int]int
var tozi []int
var n int

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	arr := strings.Split(string(line), " ")

	n, _ = strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	k, _ := strconv.Atoi(arr[2])

	slip = make(map[int]int)
	tozi = make([]int, m, m)

	for i := 0; i < m; i++ {
		line, _, _ := reader.ReadLine()
		key, _ := strconv.Atoi(string(line))
		setData(i, key)
	}

	for i := 0; i < k; i++ {
		for key, _ := range tozi {
			next := nextWeizhi(key)
			setData(key, next)
		}
	}

	for _, value := range tozi {
		fmt.Println(value)
	}

}

func nextWeizhi(tozhikey int) int {
	weizhiKey := tozi[tozhikey]
	return noExTozhi(weizhiKey)
}

func noExTozhi(weizhiKey int) int {
	if _, isEx := slip[weizhiKey]; isEx {
		if weizhiKey++; weizhiKey > n {
			weizhiKey = 1
		}
		return noExTozhi(weizhiKey)
	} else {
		return weizhiKey
	}
}

func setData(tozikey, weizhi int) {

	if tozi[tozikey] != 0 {
		delete(slip, tozi[tozikey])
	}

	slip[weizhi] = tozikey
	tozi[tozikey] = weizhi
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var arr = *new([]int)

func main() {
	runtime.GOMAXPROCS(2)

	conChan := make(chan int, itemCon)

	sc.Scan()
	itemCon, _ := strconv.Atoi(sc.Text())

	s := make([]int, itemCon)

	for k := 0; k < itemCon; k++ {
		sc.Scan()
		item, _ := strconv.Atoi(sc.Text())
		s[k] = item

	}
}

func tmp1(s []int, t1, t2 int) {

}

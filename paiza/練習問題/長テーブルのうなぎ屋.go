package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := 0
	m := 0

	line, _, _ := reader.ReadLine()

	constr := string(line)

	bs := strings.Split(constr, " ")

	n, _ = strconv.Atoi(bs[0])
	m, _ = strconv.Atoi(bs[1])

	array := make(map[int][]int)

	j := 0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		constr := string(line)

		if constr == "\\q" {
			break
		}

		bs := strings.Split(constr, " ")
		i, _ := strconv.Atoi(bs[0])
		z, _ := strconv.Atoi(bs[1])
		z--
		array[j] = []int{i, z}
		j++
	}

	//椅子を真似する
	yisi := make([]int, n, n)

	str := ""

	for t := 0; t < m; t++ {
		i, z := array[t][0], array[t][1]
		str += string(i) + " " + string(z)
		if z <= n-1 {
			if i+z <= n {
				if isAdd(yisi, z, i+z) {
					addYisi(&yisi, z, i+z, t+1)
				}

			} else {
				if isAdd(yisi, z, n) && isAdd(yisi, 0, i-n+z) {
					addYisi(&yisi, z, n, t+1)
					addYisi(&yisi, 0, i-n+z, t+1)
				}

			}
		}
	}

	cont := 0
	for _, tmp2 := range yisi {
		if tmp2 != 0 {
			cont++
		}
	}

	fmt.Println(cont)

}

func isAdd(yisi []int, z, i int) bool {

	tmp1 := yisi[z:i]
	isAdd := true
	for _, tmp2 := range tmp1 {
		if tmp2 != 0 {
			isAdd = false
			break
		}
	}

	return isAdd
}

func addYisi(yisiP *[]int, z, i, su int) {
	yisi := *yisiP
	tmp1 := yisi[z:i]
	for key, _ := range tmp1 {
		tmp1[key] = su
	}

}

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	alln := big.NewInt(1)

	n := 2

	for {
		alln.Mul(alln, big.NewInt(int64(n)))

		t := get(alln)
		t2 := add(n)
		if t != t2 {
			fmt.Println(t, t2, n)
			break
		}
		n++
		if n == 1000001 {
			break
		}
	}
}

func get(allns *big.Int) int {

	alln := big.NewInt(1)
	alln.Mul(alln, allns)

	allStr := alln.String()

	//fmt.Println(allStr)

	lens := len(allStr)
	no0 := lens - 1

	for i := lens - 1; i >= 0; i-- {

		if allStr[i] != 48 {
			no0 = i
			break
		}
	}

	allStr = allStr[0 : no0+1]

	//fmt.Println(allStr)

	lens = len(allStr)

	if lens >= 9 {
		allStr = allStr[lens-9 : lens]
	}

	lens = len(allStr)

	//fmt.Println("----------")
	//fmt.Println(allStr)

	for i := 0; i < lens; i++ {
		if allStr[i] != 48 {
			no0 = i
			break
		}
	}

	reN, _ := strconv.Atoi(allStr[no0:lens])

	return reN
}

func add(n int) int {

	var m int64
	m = 1

	for {
		m *= int64(n)

		if m%10 == 0 {

			for {
				if m%10 == 0 {
					m /= 10
				} else {
					break
				}
			}
		}
		m = m % 100000000000
		n--
		if n == 1 {
			break
		}
	}
	m = m % 1000000000
	return int(m)
}

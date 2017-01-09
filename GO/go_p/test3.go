package main

import (
	"fmt"
	"runtime"
)

var ss = map[int]int{}

func main() {

	runtime.GOMAXPROCS(2)
	itemCon := 11

	for kk := 0; kk < itemCon; kk++ {
		t := kk % 7
		if t < 7 {
			if _, ok := ss[t]; ok {
				ss[t]++
			} else {
				ss[t] = 1
			}
		}
		fmt.Println(kk)
	}

	arr := []int{0, 1, 2, 3, 4, 5, 6}
	con := 0

	for z := 0; z < 7-2; z++ {
		for j := z + 1; j < 7-1; j++ {
			con += tmp1(arr, z, j)
		}
	}

	fmt.Println(ss)
	fmt.Println(con)

}

func tmp1(s []int, i, k int) int {
	con := 0

	key1 := s[i]
	key2 := s[k]
	for _, item := range s[k+1:] {
		if ((key1 + key2 + item) > 7) && (key1+key2+item)%7 == 0 {
			con += ss[key1] * ss[key2] * item
		}
	}

	return con
}

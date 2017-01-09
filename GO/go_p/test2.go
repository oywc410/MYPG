package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(2)
	itemCon := 10000

	s := make([]int, itemCon)

	for kk := 0; kk < itemCon; kk++ {
		s[kk] = kk
	}

	conChan := make(chan int, 1000)

	t := 0

	for z := 0; z < itemCon-2; z++ {

		t++
		j := z + 1

		go func(s []int, z, j int) {
			con := 0
			for ; j < itemCon-1; j++ {

				if j%3 == 0 {
					runtime.Gosched()
				}

				con += tmp1(s, z, j)
			}
			fmt.Println(con)
			conChan <- con
			runtime.Goexit()
		}(s, z, j)
	}

	con := 0
LOOP:
	for {
		select {
		case c := <-conChan:
			con += c
			t--
			if t <= 0 {
				break LOOP
			}
		}
	}

	fmt.Println(con)

}

func tmp1(s []int, i, k int) int {
	con := 0

	key1 := s[i]
	key2 := s[k]
	for _, item := range s[k+1:] {
		is := (key1 + key2 + item) % 7

		if is == 0 {
			con++
		}
	}

	return con
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
)

func main() {

	runtime.GOMAXPROCS(2)

	reader := bufio.NewReader(os.Stdin)

	itemCon := 0
	var s []int

	k := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if itemCon == 0 {

			itemCon, _ = int(string(line))
			s = make([]int, itemCon)
		} else {
			s[k], _ = int(string(line))
			k++
		}
		if k == itemCon {
			break
		}
	}

	conChan := make(chan int, 100)

	t := 0

	for z := 0; z < itemCon-2; z++ {

		t++
		j := z + 1

		go func(s []int, z, j int) {
			con := 0
			for ; j < itemCon-1; j++ {
				con += tmp1(s, z, j)
			}
			conChan <- con
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
		if ((key1 + key2 + item) > 7) && (key1+key2+item)%7 == 0 {
			con++
		}
	}

	return con
}

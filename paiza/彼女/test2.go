package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	main1()
}

func main1() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	n, _ := strconv.Atoi(string(line))

	add(n)
}

func main2() {
	for i := 2; i <= 30; i++ {
		add(i)
	}
}

func add(n int) {

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
		m = m % 1000000000000
		n--
		if n == 1 {
			break
		}
	}
	m = m % 1000000000
	fmt.Println(m)
}

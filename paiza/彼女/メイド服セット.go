package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()
	count, _ := strconv.Atoi(string(line))

	item := 0
	var itemData []int = make([]int, count)

	for {
		if count == item {
			break
		}

		line, _, _ := reader.ReadLine()
		itemData[item], _ = strconv.Atoi(string(line))
		item++
	}

	for _, value := range itemData  {

		if value == 0 {
			fmt.Println("01:00")
			continue
		}

		time := value / 3
		h := 25 - time / 60

		var i int

		if time != 0 {
			i = 60 - time % 60
			if i > 0 {
				h--
			}

			if i == 60 {
				h++
				i = 0
			}
		}

		if h > 23 {
			h = h % 24
		}

		fmt.Printf("%02d:%02d\n", h, i)
	}
}

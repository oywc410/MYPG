package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	topStr := strings.Split(string(line), " ")

	x, _ := strconv.Atoi(topStr[0])
	y, _ := strconv.Atoi(topStr[1])
	z, _ := strconv.Atoi(topStr[2])
	catCon, _ := strconv.Atoi(topStr[3])

	catX := make([]int, 0, catCon)
	catY := make([]int, 0, catCon)

	for {
		line, _, _ = reader.ReadLine()
		catStr := strings.Split(string(line), " ")
		catI, _ := strconv.Atoi(catStr[1])

		if catStr[0] == "0" {
			catX = append(catX, catI)
		} else {
			catY = append(catY, catI)
		}

		catCon--
		if catCon == 0 {
			break
		}
	}

	catX = append(catX, x)
	catY = append(catY, y)

	sort.Ints(catX)
	sort.Ints(catY)

	var catXLin []int
	var catYLin []int

	tmpX := 0

	for _, value := range catX {
		catXLin = append(catXLin, value-tmpX)
		tmpX = value
	}

	tmpY := 0

	for _, value := range catY {
		catYLin = append(catYLin, value-tmpY)
		tmpY = value
	}

	min := 0

	for _, vX := range catXLin {
		for _, vY := range catYLin {
			if min == 0 {
				min = vX * vY
			} else {
				tmp := vX * vY
				if min > tmp {
					min = tmp
				}
			}
		}
	}

	fmt.Println(min * z)
}

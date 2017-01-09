package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	str := string(line)
	strSilp := map[string]int{
		"c": 0,
		"a": 0,
		"t": 0,
	}

	for _, st := range str {
		strSilp[string(st)]++
	}

	isCheck := false
	var minVal int
	var maxVal int

	for _, val := range strSilp {
		if isCheck == false {
			minVal = val
			isCheck = true
			maxVal = val
		} else {
			if minVal > val {
				minVal = val
			}
			if maxVal < val {
				maxVal = val
			}
		}
	}

	if minVal == maxVal && minVal == 0 {
		fmt.Println(0)
		fmt.Println(0)
		fmt.Println(0)
		fmt.Println(0)
	} else {
		fmt.Println(minVal)
		maxVal -= minVal
		strSilp["c"] -= minVal
		strSilp["a"] -= minVal
		strSilp["t"] -= minVal
		fmt.Println(maxVal - strSilp["c"])
		fmt.Println(maxVal - strSilp["a"])
		fmt.Println(maxVal - strSilp["t"])
	}
}

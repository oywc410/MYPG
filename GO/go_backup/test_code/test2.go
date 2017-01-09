package main

import (
	"fmt"
)

func main() {

	str := "\\asdasd\\as\\d\\asdg\\ds\\g\\sdfasdfsdf"

	fmt.Println(str)
	fmt.Println(catNewDir(str))
}

func catNewDir(str string) string {
	rs := []rune(str)
	end := 0
	for keys, a := range rs {
		if a == 92 {
			end = keys
		}
	}

	return string(rs[:end])
}

package main

import (
	"fmt"
	"sort"
)

type StringSlicep []string

func (p StringSlicep) Len() int {
	return len(p)
}

func (p StringSlicep) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlicep) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	names := make([]string, 5)
	names[0] = "aaaa"
	names[1] = "bbbb"

	sort.Sort(StringSlicep(names))
	fmt.Println(names)
}

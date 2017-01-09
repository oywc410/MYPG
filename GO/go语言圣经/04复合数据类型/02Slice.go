package main
import (
	"fmt"
	"unicode"
)

func main() {
	s := []int{5, 6, 7, 8}
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	fmt.Println(unicode.IsSpace(' '))
	return slice[:len(slice)-1]
}

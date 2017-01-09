package main
import (
	"fmt"
	"unicode/utf8"
	"strings"
	"bytes"
)

func main() {

	//字符串长度
	s := "Hello, 世界"

	fmt.Println("len:", len(s))
	fmt.Println("utf8 len:", utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d\t%q\n", s[i], string(s[i]))
	}

	//获取文件名
	fmt.Println(basename("a/b/filename.go"))

	fmt.Println(intsToString([]int{1, 2, 3}))

}

func basename(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

//简化版
func basename02(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}

		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}


/**
bytes strings 实用函数
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
 */
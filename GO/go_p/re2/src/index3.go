package main

import (
	"fmt"
	"regexp"
)

func main() {

	//查找函数
	/**
	func (re *Regexp) Find(b []byte) []byte
	func (re *Regexp) FindAll(b []byte, n int) [][]byte
	func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
	func (re *Regexp) FindAllString(s string, n int) []string
	func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
	func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
	func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
	func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
	func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
	func (re *Regexp) FindIndex(b []byte) (loc []int)
	func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
	func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
	func (re *Regexp) FindString(s string) string
	func (re *Regexp) FindStringIndex(s string) (loc []int)
	func (re *Regexp) FindStringSubmatch(s string) []string
	func (re *Regexp) FindStringSubmatchIndex(s string) []int
	func (re *Regexp) FindSubmatch(b []byte) [][]byte
	func (re *Regexp) FindSubmatchIndex(b []byte) []int
	*/

	a := "I am learning Go launguage"

	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串,不然就是放回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)laun(.*)")
	//查找submantch,返回数组,第一个元素的匹配的全部元素,第二元素师第一个()里面的....
	//下面的输出第一个元素是"am learning go language"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	submatchall := re2.FindAllSubmatch([]byte(a), 1)
	fmt.Println(submatchall)

	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

	//替换函数
	/**
	func (re *Regexp) ReplaceAll(src, repl []byte) []byte
	func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
	func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
	func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
	func (re *Regexp) ReplaceAllString(src, repl string) string
	func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
	*/
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)

	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}

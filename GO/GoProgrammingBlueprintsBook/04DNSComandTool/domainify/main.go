package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"strings"
	"unicode"
	"fmt"
)

var tlds = []string{"com", "net"}
//允许使用的字符串
const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"
//生成完整域名
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			//空格转[-]
			if unicode.IsSpace(r) {
				r = '-'
			}
			//过滤字符串
			// ContainsRune 判断字符串 s 中是否包含字符
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}

//echo "aaa" | ..\sprinkle\sprinkle.exe | .\domainify.exe
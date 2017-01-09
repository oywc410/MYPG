package main

import (
	"fmt"
	"os"
	"regexp"
)

/**
使用正则来匹配验证内容

func Match(pattern string, b []byte) (matched bool, error error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, error error)
func MatchString(pattern string, s string) (matched bool, error error)
上面的三个函数实现了同一个功能，就是判断pattern是否和输入源匹配，匹配的话就返回true，如果解析正则
出错则返回error。三个函数的输入源分别是byte slice、RuneReader和string。
*/

func IsIp(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip); !m {
		return false
	}

	return true
}

func main() {
	//接受用户输入
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
}

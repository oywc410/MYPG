package main

import (
	"fmt"
	"text/template"
)

/**
模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变
量是否正确的书写。
*/

func main() {
	tOk := template.New("fires")
	template.Must(tOk.Parse("some static text /* and a comment */"))
	fmt.Println("The first one parsed ok.")

	template.Must(template.New("secaond").Parse("some static text {{ .Name }}"))
	fmt.Println("The first one parsed ok.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}

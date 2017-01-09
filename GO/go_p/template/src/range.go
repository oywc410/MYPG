package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

/**
{{range}} 这个和Go语法里面的range类似，循环操作数据
{{with}}操作是指当前对象的值，类似上下文的概念
*/

func main() {
	f1 := Friend{Fname: "mianx.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("filedname example")
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
		{{end}}
		`)
	p := Person{UserName: "ABCNAME",
		Emails:  []string{"emial@sf.com", "asda@gasd.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

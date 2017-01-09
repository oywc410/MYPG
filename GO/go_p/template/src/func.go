package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	//find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + " at " + substrs[1])
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("filedname example")

	//为模板添加函数
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an emails {{. |emailDeal}}
			{{end}}
			{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
			{{end}}
		`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@ERFSAF.COM", "ASDAS@ESAF.COM"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

/**
GO模板中自带的函数:
var builtins = FuncMap{
	"and": and,
	"call": call,
	"html": HTMLEscaper,
	"index": index,
	"js": JSEscaper,
	"not": not,
	"or": or,
	"print": fmt.Sprint,
	"printf": fmt.Sprintf,
	"println": fmt.Sprintln,
	"urlquery": URLQueryEscaper,
}
*/

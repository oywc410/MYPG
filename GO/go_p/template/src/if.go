package main

import (
	"os"
	"text/template"
)

type HtmlTplP struct {
	Htmls string
}

func main() {
	tEmpty := template.New("template test")
	//Must函数用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化：
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo : {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)
	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
	//注意：if里面无法使用条件判断，例如.Mail==“astaxie@gmail.com”，这样的判断是不正确的，if里面只能是bool值

	mm := HtmlTplP{Htmls: "<a>test</a>"}
	tHtmlValue := template.New("templae test")
	tHtmlValue = template.Must(tHtmlValue.Parse("{{.Htmls | html}}"))
	tHtmlValue.Execute(os.Stdout, mm)

	//Go语言模板最强大的一点就是支持pipe数据，在Go语言里面任何{{}}里面的都是pipelines数据，例如我们上面输出的email里面如果还有一些可能引起XSS注入的，那么我们如何来进行转化呢？
	//{{. | html}}在emai

	tPValue := template.New("template test")
	tPValue = template.Must(tPValue.Parse(`
			{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
			{{with $x := "output"}}{{printf "%q" $x}}{{end}}
			{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
		`))
	tPValue.Execute(os.Stdout, nil)
	//模板变量
}

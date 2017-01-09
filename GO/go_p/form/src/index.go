package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"
	"regexp"
	"time"
	"crypto/md5"
	"io"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	//解析URL传递的参数,对于POST则解析响应包的主题(request body)
	//注意:如果没有调用ParseForm方法，下面无法获取表单数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //写入掉客户端
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)//获取请求方法
	if r.Method == "GET" {
	
		t := template.Must(template.ParseFiles("login.gtpl"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := t.Execute(w, nil); err != nil {
			fmt.Println(err)
		}
	} else {
		//请求的是登陆数据,那么执行登陆逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		
		//request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息，下面展示了可以对form数据
		//进行的一些操作:
		v := r.Form
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("friend"))
		fmt.Println(v["friend"])
		
		
		//表单处理
		if len(r.Form["username"][0]) == 0 {
		}
		
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			//数字转化错误,那么可能就不是数字
		}
		
		if getint > 100 {
			//太大了
		}
		
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			//正则匹配
		}
		
		//判断是否为中午
		if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m{
		
		}
		
		//判断是否英文
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		
		}
		
		//邮件号码
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}
		
		
		//手机号码
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		
		}
		
		//数组存在判断
		/*
		slice:=[]string{"apple","pear","banane"}
		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				return true
			}
		}
		return false
		*/
		//时间
		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())
		
		//身份证
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		
		}
		
		
		//XSS
		/*
		func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
		func HTMLEscapeString(s string) string //转义s之后返回结果字符串
		func HTMLEscaper(args …interface{}) string //支持多个参数一起转义，返回结果字符串
		*/
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		
		//模板中的变量会自行转义 为防止使用template.HTML
		
		/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(out, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		*/
		
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		fmt.Println("-----")
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t := template.Must(template.ParseFiles("upload.gtpl"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, token)
		
	} else {
		r.ParseMultipartForm(32 << 20)//32 << 20上传文件储存在内存里面,如果超过大小前储存在系统的临时文件夹中
		file, handler, err := r.FormFile("uploadfile")//获取上传文件句柄
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()//发生错误时 调用
		fmt.Fprintf(w, "%v", handler.Header)
		//新建文件 读写 权限 0666
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()//发生错误时调用
		io.Copy(f, file)//写入文件
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)//设置访问的路由
	http.HandleFunc("/login", login)//设置访问的路由
	http.HandleFunc("/upload", upload)//文件上传
	err := http.ListenAndServe(":9090", nil)//设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
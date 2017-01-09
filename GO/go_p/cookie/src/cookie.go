package main

// 代码包导入语句。
import (
	"fmt" // 导入代码包fmt。
	"log"
	"net/http"
	"time"
)

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

func action(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration.AddDate(1, 0, 0)
	//写入cookie
	cookie := http.Cookie{Name: "user", Value: "value", Expires: expiration, MaxAge: 16}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w, cookie.Value)
}

func action2(w http.ResponseWriter, r *http.Request) {

	//读取cookie
	cookie, _ := r.Cookie("user")

	fmt.Fprint(w, cookie)

	fmt.Println("cookie:", cookie)

	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}

func main() {

	http.HandleFunc("/", action) //设置访问的路由
	http.HandleFunc("/get", action2)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

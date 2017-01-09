package main

import (
	"../trace"
	"flag"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"github.com/stretchr/objx"
)

var avatars Avatar = TryAvatars{
	UserFileSystemAvatar,
	UseAuthAvatar,
	UserGravatar,
}

func main() {

	var addr = flag.String("addr", ":8001", "アプリケーションのアドレス")
	flag.Parse()

	gomniauth.SetSecurityKey("mysrcuritykey")
	gomniauth.WithProviders(
		//facebook.New("???", "???", "http://......")
		google.New("400720597353-vb6mkg2ia79blui2qs2vbc67dedtnrns.apps.googleusercontent.com", "qUDc1kADhKC7dGjdNNI1MVU7", "http://127.0.0.1:8001/auth/callback/google"),
	)

	//r := newRoom(UseAuthAvatar)
	//r := newRoom(UserGravatar)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	//StripPrefix显示文件目录列表
	//FileServer文件服务器
	http.Handle("/avatars/", http.StripPrefix("/avatars", http.FileServer(http.Dir("./avatars"))))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	//http.Handler("/assets/", http.FileServer(http.Dir("./assets")))

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)

	// '/root'  只解析到/root       '/root/'  解析/root/XXXXX 至结束

	//启动socket监听
	go r.run()

	log.Println("Web服务器启动,端口:", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

type templateHandler struct {
	once     sync.Once //go中的单例模式的实现
	filename string
	temp1    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	t.once.Do(func() { //此处代码只执行一次 多个携程中也只执行一次
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("tpl", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	if err := t.temp1.Execute(w, data); err != nil {
		fmt.Println(err)
	}
}

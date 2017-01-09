package main
import (
	"net/http"
	"strings"
	"log"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"crypto/md5"
	"io"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		//未认证
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		//错误
		panic(err.Error())
	} else {
		//成功
		h.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//地址类型 /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login"://登入
		log.Println("TODO: 登入处理", provider)
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("认证获取失败", provider, "-", err)
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("GetBeginAuthUrl获取失败", provider, "-", err)
		}
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback"://认证成功返回信息解析
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("认证获取失败", provider, "-", err)
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("认证中断未完成", provider, "-", err)
		}

		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("用户信息获取失败", provider, "-", err)
		}

		chatUser := &chatUser{User: user}
		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Name()))
		chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
		avatarURL, err := avatars.GetAvatarURL(chatUser)
		if err != nil {
			log.Fatalln("GetAvatarURL获取失败", "-", err)
		}

		authCookieValue := objx.New(map[string]interface{}{
			"userid": chatUser.uniqueID,
			"name": user.Name(),
			"avatar_url": avatarURL,
		}).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: authCookieValue,
			Path: "/",
		})

		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "操作%s非对应", action)
	}
}

type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}
package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"sync"
	"net/url"
	"html/template"
	"time"
)

//Provider接口 (用户sessuib的数据操作)
type Provider interface {
	//初始化session
	SessionInit(sid string) (Session, error)
	//返回所要的session
	SessionRead(sid string) (Session, error)
	//销毁session
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

//Session接口
type Session interface {
	Set(key, value interface{}) error
	get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

//session对象
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

var globalSessions *session.Manager
var provides = make(map[string]Provide)

func Register(name string, provide Provide) {
	if driver == nil {
		panic("session: Register provide Provide")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide" + name)
	}
	provides[name] = provide
}

func (Manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value = "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

//session重置
func (mananger *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		mananger.lock.Lock()
		defer mananger.lock.Unlock()
		mananger.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager,GC() })
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provider[provideName]
	if !ok {
		return nil, fmt.Errorf("session : unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

func init() {
	go globalSessions.GC()
}

func main() {
	globalSessions = NewManager("memory", "gosessionid", 3600)
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
	}
}

//计数器
func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, y)
		sess = globalSessions.SessionStart(w, r)
	}

	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}

	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

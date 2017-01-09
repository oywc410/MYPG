package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"flag"
	"log"
	"time"
	"github.com/stretchr/graceful"
)

func main() {
	var (
		addr = flag.String("addr", ":8005", "接听地址")
		mongo = flag.String("mongo", "localhost", "MongoDB地址")
	)

	flag.Parse()
	log.Println("MaongoDB链接中")
	db, err := mgo.Dial(*mongo)
	if err != nil {
		log.Fatalf("MaongoDB链接失败:", err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/polls/", withCORS(withVars(withData(db, withAPIKey(handlePolls)))))
	log.Println("web服务器启动中:", *addr)
	graceful.Run(*addr, 1 * time.Second, mux)
	log.Println("停止中")
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidAPIKey(r.URL.Query().Get("key")) {
			respondErr(w, r, http.StatusUnauthorized, "API KEY ERROR")
			return
		}
		fn(w, r)
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}

//全局(当前链接有效)数据库链接
func withData(d *mgo.Session, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thisDb := d.Copy()
		defer thisDb.Close()
		//每个链接保持DB SESSION
		SetVar(r, "db", thisDb.DB("ballots"))
		f(w, r)
	}
}

//全局(当前链接有效)变量
func withVars(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		OpenVars(r)
		defer CloseVars(r)
		fn(w, r)
	}
}

//AJAX访问 夸域名对策
func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}

//解析JSON数据
func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

//编码JSON数据
func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

//输出常用http错误信息
func respondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	respondErr(w, r, status, http.StatusText(status))
}
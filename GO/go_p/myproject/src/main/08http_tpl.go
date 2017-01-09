package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// パラメータを取得
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}
	filename := fmt.Sprintf("%d.txt", id)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// personを生成
	person := Person{
		ID:   id,
		Name: string(b),
	}
	// レスポンスにエンコーディングしたHTMLを書き込む
	t.Execute(w, person)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3001", nil)
}

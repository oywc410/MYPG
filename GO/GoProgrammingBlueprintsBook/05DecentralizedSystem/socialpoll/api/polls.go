package main

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"gopkg.in/mgo.v2"
)

type poll struct {
	ID bson.ObjectId `bson:"_id" json:"id"` // bson对象 ID名为_id  json对象 ID名为id
	Title string `json:"title"`
	Options []string `json:"options"`
	Results map[string]int `json:"results,omitempty"`
}

func handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlePollsGet(w, r)
		return
	case "POST":
		handlePollsPost(w, r)
		return
	case "DELETE":
		handlePollsDelete(w, r)
		return
	case "OPTIONS":
		w.Header().Add("Access-Control-Allow-Methods", "DELETE")
		respond(w, r, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, r, http.StatusFound)
}

func handlePollsGet(w http.ResponseWriter, r *http.Request) {

	db := GetVar(r, "db").(*mgo.Database)
	c := db.C("polls")
	var q *mgo.Query
	p := NewPath(r.URL.Path)
	if p.HasID() {
		//特定查询
		q = c.FindId(bson.ObjectIdHex(p.ID))
	} else {
		//查询所有
		q = c.Find(nil)
	}

	var result []*poll
	if err := q.All(&result); err != nil {
		respondErr(w, r, http.StatusInternalServerError, err)
		return
	}
	respond(w, r, http.StatusOK, &result)
}

func handlePollsPost(w http.ResponseWriter, r *http.Request) {
	db := GetVar(r, "db").(*mgo.Database)
	c := db.C("polls")
	var p poll
	if err := decodeBody(r, &p); err != nil {
		respondErr(w, r, http.StatusBadRequest, "请求失败:", err)
		return
	}
	p.ID = bson.NewObjectId()
	if err := c.Insert(p); err != nil {
		respondErr(w, r, http.StatusInternalServerError, "数据保存失败", err)
		return c
	}

	w.Header().Set("Location", "polls/" + p.ID.Hex())
	respond(w, r, http.StatusCreated, nil)
}

func handlePollsDelete(w http.ResponseWriter, r *http.Request) {
	db := GetVar(r, "db").(*mgo.Database)
	c := db.C("polls")
	p := NewPath(r.URL.Path)
	if !p.HasID() {
		respondErr(w, r, http.StatusMethodNotAllowed, "禁止删除所有项目")
		return
	}
	if err := c.RemoveId(bson.ObjectIdHex(p.ID)); err != nil {
		respondErr(w, r, http.StatusInternalServerError, "数据删除失败:", err)
		return
	}
	respondErr(w, r, http.StatusOK, nil)
}

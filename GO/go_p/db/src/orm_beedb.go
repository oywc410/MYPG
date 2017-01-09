package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

type Userinfo struct {
	Uid int `PK`
	Username string
	Departname string
	Created time.Time
}

func main() {
	db, err := sql.Open("mymysql", "go_test/root/root")
	if err != nil {
		panic(err)
	}
	
	beedb.OnDebug = true
	orm := beedb.New(db)
	
	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)
	
	add := make(map[string]interface{})
	add["username"] = "astaxie"
	add["departname"] = "cloud develop"
	add["created"] = "2012-12-12"
	orm.SetTable("userinfo").Insert(add)
	
	addslice := make([]map[string]interface{})
	add := make(map[string]interface{})
	add2 := make(map[string]interface{})
	add["username"] = "astaxie"
	add["departname"] = "cloud develop"
	add["created"] = "2020-12-22"
	add2["username"] = "astaxie2"
	add2["departname"] = "cloud develop2"
	add2["created"] = "2012-02-02"
	addslice = append(addslice, add, add2)
	orm.SetTable("userinfo").Insert(addslice)
}

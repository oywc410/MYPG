package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)


/**
https://github.com/bmizerany/pq 支持database/sql驱动，纯Go写的
https://github.com/jbarham/gopgsqldriver 支持database/sql驱动，纯Go写的
https://github.com/lxn/go-pgsql 支持database/sql驱动，纯Go写的
*/

func main() {
	db, err := sql.Open("postgres", "user=postgres password=root dbname=go_test sslmode=disable")
	checkErr(err)
	
	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo (username, departname, created) VALUES ($1, $2, $3) RETURNING uid")
	checkErr(err)
	
	res, err := stmt.Exec("astaxie", "研发部门", "2012-11-11")
	checkErr(err)
	
	//pg不支持这个函数,因为它没有类似mysql的自增ID
	id, err := res.LastInsertId()
	checkErr(err)
	
	fmt.Println(id)
	
	//更新数据
	stmt, err = db.Prepare("update userinfo set username = $1 where uid = $2")
	checkErr(err)
	
	res, err = stmt.Exec("astaxieupdate", 1)
	checkErr(err)
	
	affect, err := res.RowsAffected()
	checkErr(err)
	
	fmt.Println(affect)
	
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
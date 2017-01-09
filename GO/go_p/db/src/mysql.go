package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"//不自动调用init
	"database/sql"
	"fmt"
	//"time"
)

func main() {
	/**
		打开数据库连接
		user@unix(/path/to/socket)/dbname?charset=utf8
		user:password@tcp(localhost:5555)/dbname?charset=utf8
		user:password@/dbname
		user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
		
		常用驱动
		https://github.com/Go-SQL-Driver/MySQL 支持database/sql，全部采用go写。
		https://github.com/ziutek/mymysql 支持database/sql，也支持自定义的接口，全部采用go写。
		https://github.com/Philio/GoMySQL 不支持database/sql，自定义接口，全部采用go写。
	*/
	db, err := sql.Open("mysql", "root:root@/go_test")
	checkErr(err)
	
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	
	res, err := stmt.Exec("astaxie", "test", "2012-12-09")
	checkErr(err)
	
	res, err = stmt.Exec("astaxie", "test2", "2012-12-09")
	checkErr(err)
	
	//获取最后插入的ID
	id, err := res.LastInsertId()
	checkErr(err)
	
	fmt.Println(id)
	
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	
	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)
	
	//影响条数
	affect, err := res.RowsAffected()
	checkErr(err)
	
	fmt.Println(affect)
	
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	
	//输出查询值
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		//传入需要赋值变量的地址
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	
	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid = ?")
	checkErr(err)
	
	res, err = stmt.Exec(id)
	checkErr(err)
	
	fmt.Println(affect)
	
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

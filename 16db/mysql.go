package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/go_test?charset=utf8")
	checkErr(err)

	defer db.Close()

	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("simple", "development", time.Now())
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("res id => ", id)

	stmt, err = db.Prepare("update userinfo set username = ? where uid = ?")
	checkErr(err)
	res, err = stmt.Exec("update new", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("affect => ", affect)

	rows, err := db.Query("select * from userinfo")
	checkErr(err)
	fmt.Println("uid\tusername\tdepartment\tcreated")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, "\t", username, "\t", department, "\t", created)
	}

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("delete affect => ", affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

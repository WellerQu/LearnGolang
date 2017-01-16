package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "root:12345678@/nixondb?charset=utf8")
	checkErr(err)

	stmt, err := conn.Prepare("insert into userinfo (uname, created) values (?, ?)")
	checkErr(err)

	res, err := stmt.Exec("cat", "20170115")
	checkErr(err)

	fmt.Println(res)

	defer func() {
		fmt.Println("close connection")
		conn.Close()
	}()

	fmt.Println("vim-go")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

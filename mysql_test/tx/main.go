package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	//root:password@tcp(ip:port)/database
	database, err :=
		sqlx.Open("mysql", "root:123456@tcp(101.37.20.199:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed:,", err)
		return
	}
	res, err :=
		conn.Exec("insert into people(username, sex, email) values (?,?,?)", "stud001", "man", "stu001@qq.com")
	if err != nil {
		fmt.Println("exec failed,", err)
		conn.Rollback()
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("get id failed,", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success:id=", id)
	res, err = conn.Exec("insert into people(username, sex, email) VALUES (?,?,?)", "sut002", "woman", "sut002@qq.com")
	if err != nil {
		fmt.Println("exec failed,", err)
		conn.Rollback()
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		fmt.Println("get id failed:id=", id)
		conn.Rollback()
		return
	}
	fmt.Println("insert success:id=", id)
	conn.Commit()
}

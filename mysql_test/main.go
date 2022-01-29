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
	//1. insert
	result, err :=
		Db.Exec("insert into person(username,sex,email) values(?,?,?)", "stu002", "woman", "stu02@qq.com")
	if err != nil {
		fmt.Println("insert failed,", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("insert success: id=", id)
	//2. select
	var person []Person
	err =
		Db.Select(&person, "select * from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("select success:person=", person)
	//3. update
	result, err = Db.Exec("update person set username=? where user_id=?", "stu0001", 2)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
		return
	}
	fmt.Println("update success,row=", row)
	//4. delete
	res, err := Db.Exec("delete from person where user_id=?", 3)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err = res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
		return
	}
	fmt.Println("delete success:row=", row)
	defer Db.Close()
}

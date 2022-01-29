package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	UserId int `gorm:"primary_key"`
	//default标签没用
	Username string `gorm:"default:'park'"`
	Sex      string
	Email    string
	Age      int
}

type Place struct {
	Country string
	City    string
	TelCode int
}

func main() {
	db, err :=
		gorm.Open("mysql", "root:123456@tcp(101.37.20.199:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed establish connection,err:", err)
		return
	}
	var person1 Person
	var person2 Person
	var person3 Person
	var person4 Person

	//获取第一条记录，按照主键排序
	////select * from people order by id limit 1
	db.First(&person1)
	fmt.Println("first person=", person1)

	//获取一条记录，不指定排序
	////select * from people limit 1
	db.Take(&person2)
	fmt.Println("person=", person2)

	//获取最后一条记录，不指定排序
	////select * from people order by id desc limit 1
	db.Last(&person3)
	fmt.Println("last person=", person3)

	var people []Person
	//获取所有记录
	//select * from people
	db.Find(&people)
	fmt.Println("people=", people)

	//通过主键查询
	////select * from users where id = 10
	db.First(&person4, 10)
	fmt.Println("person=", person4)

	var person5 Person

	var people2 []Person
	var people3 []Person
	var people4 []Person
	var people5 []Person

	//获取第一条匹配的记录
	////select * from people where name ='Lucy' limit 1
	db.Where("username=?", "Lucy").First(&person5)
	fmt.Println("username=Lucy,person=", person5)

	//获取所有匹配的记录
	////select * from people where sex='Boy'
	db.Where("sex=?", "Boy").Find(&people2)
	fmt.Println("sex=Boy,people=", people2)

	//<>
	////select * from people where name <> 'Lucy'
	db.Where("username <> ?", "Lucy").Find(&people3)
	fmt.Println("username <> Lucy,people=", people3)

	//And
	////select * from people where name = 'Lucy' and sex <> 'Boy'
	db.Where("username = ? and sex <> ?", "Lucy", "Boy").Find(&people4)
	fmt.Println("username = Lucy and sex <> Boy,people=", people4)

	//Struct
	//// select * from people where username ='Jack and sex ='Boy'
	db.Where(Person{Username: "Jack", Sex: "Boy"}).Find(&people5)
	fmt.Println("username = jack and sex = Boy,people=", people5)

	defer db.Close()
}

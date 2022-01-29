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

	park := Person{
		UserId:   23,
		Username: "Park",
		Sex:      "Boy",
		Email:    "Park@gmail.com",
		Age:      100,
	}
	park.Sex = "Girl"
	park.Age = 98
	//会更新所有的字段，即使这些字段没有被修改
	saved := db.Save(&park)
	fmt.Println("error=", saved.Error)

	updated := db.Model(&park).Update("sex", "Boy")
	fmt.Println("error=", updated.Error)

	defer db.Close()
}

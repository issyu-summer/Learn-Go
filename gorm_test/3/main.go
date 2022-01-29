package main

import (
	"errors"
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

func (p *Person) BeforeCreate(scope *gorm.Scope) (err error) {
	//不管用
	//scope.Set("age", 888)
	//用来进行字段的检查
	if p.Age == 0 {
		err = errors.New("age can't be 0")
		//有用，但创建会失败，所以该语句也会无效
		p.Age = 999
	}
	return
}

func main() {
	db, err :=
		gorm.Open("mysql", "root:123456@tcp(101.37.20.199:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed establish connection,err:", err)
		return
	}
	park := Person{
		Username: "Park",
		Sex:      "Boy",
		Email:    "Park@gmail.com",
		Age:      100,
	}
	//检查主键是否为空
	blankKey := db.NewRecord(park)
	fmt.Println("主键为空？", blankKey)
	//创建记录
	created := db.Create(&park)
	fmt.Println("创建成功？", created.Error)
	//检查主键是否为空
	blankKey = db.NewRecord(park)
	fmt.Println("主键为空？", blankKey)
	fmt.Println("park=", park)
	defer db.Close()
}

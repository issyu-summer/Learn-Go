package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	//设置字段的大小为255个字节
	Role string `gorm:"size:255"`
	// 设置 memberNumber 字段唯一且不为空
	MemberNumber *string `gorm:"unique;not null"`
	// 设置 Num字段自增
	Num int `gorm:"AUTO_INCREMENT"`
	// 给Address 创建一个名字是  `addr`的索引
	Address string `gorm:"index:addr"`
	//忽略这个字段
	IgnoreMe int `gorm:"-"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(101.37.20.199:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("链接成功")
	defer db.Close()
}

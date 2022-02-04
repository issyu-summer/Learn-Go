package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.创建路由
	//默认使用了2个中间件Logger()和Recovery()
	r := gin.Default()
	//路由组1，处理GET请求
	group1 := r.Group("v1")
	//{}是规范书写
	{
		group1.GET("/login", login)
		group1.GET("/submit", submit)
	}
	//路由组2，处理POST请求
	group2 := r.Group("v2")
	{
		group2.POST("/login", login)
		group2.POST("/submit", submit)
	}
	r.Run()
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, fmt.Sprintf("hello,%s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, fmt.Sprintf("hello,submit,%s\n", name))
}

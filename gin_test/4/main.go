package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		//http工具测试POST请求
		//POST http://localhost:8080/user
		//Content-Type: application/x-www-form-urlencoded
		//
		//name=summer&password=123456
		formType := c.DefaultPostForm("type", "post1")
		username := c.PostForm("name")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("name:%s,passwd:%s，type:%s", username, password, formType))
	})
	r.Run()

}

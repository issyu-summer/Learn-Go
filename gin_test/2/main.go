package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	//使用符号":"则只匹配一个，使用符号"*"则可以全部匹配
	//`https://cloud.tencent.com/developer/article/1580650`
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.Run()
}

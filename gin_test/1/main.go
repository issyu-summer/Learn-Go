package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,world")
	})
	//3.监听端口
	//Run("里面不指定端口号默认为8080")
	r.Run(":8080")
}

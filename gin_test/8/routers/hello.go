package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello,www.go.com!",
	})
}

// SetupRouter 导出的函数需要首字母大写
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	return r
}

func LoadHello(e *gin.Engine) {
	e.GET("/hello", helloHandler)
}

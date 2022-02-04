package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//使用默认值
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "summer")
		c.String(http.StatusOK, fmt.Sprintf("hello,%s", name))
	})
	r.GET("/user1", func(c *gin.Context) {
		name := c.Query("name")
		c.String(http.StatusOK, fmt.Sprintf("hello,%s", name))
	})
	r.Run()
}

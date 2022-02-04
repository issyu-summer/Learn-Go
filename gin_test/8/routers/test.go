package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test",
	})
}

func LoadTest(e *gin.Engine) {
	e.GET("/test", handleTest)
}

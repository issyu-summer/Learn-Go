package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func hello2Handler(c *gin.Context) {
	c.String(http.StatusOK, "hello2")
}

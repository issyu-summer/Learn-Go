package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testHandler(c *gin.Context) {
	c.String(http.StatusOK, "test")
}
func test2Handler(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}

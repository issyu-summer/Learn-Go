package hello

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/hello", helloHandler)
	e.GET("/hello2", hello2Handler)
}

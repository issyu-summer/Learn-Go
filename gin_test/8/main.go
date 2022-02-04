package main

import (
	"Learn-Go/gin_test/8/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//router := routers.SetupRouter()
	router := gin.Default()
	routers.LoadHello(router)
	routers.LoadTest(router)
	if err := router.Run(); err != nil {
		fmt.Println("startup service failed,err,", err)
	}
}

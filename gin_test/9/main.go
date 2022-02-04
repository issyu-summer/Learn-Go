package main

import (
	"Learn-Go/gin_test/9/hello"
	"Learn-Go/gin_test/9/routers"
	"Learn-Go/gin_test/9/test"
	"fmt"
)

func main() {
	routers.Include(test.Routers, hello.Routers)

	router := routers.Init()
	if err := router.Run(); err != nil {
		fmt.Println("startup service failed,err=", err.Error())
	}
}

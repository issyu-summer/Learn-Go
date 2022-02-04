package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.创建路由
	//默认使用了2个中间件Logger()和Recovery()
	r := gin.Default()
	//限制上传文件大小
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file failed,err=%s", err.Error()))
		}
		files := form.File["files"]
		for i, file := range files {
			if c.SaveUploadedFile(file, fmt.Sprintf("test_%d.png", i)); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload failed,err=%s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})
	r.Run(":8080")
}

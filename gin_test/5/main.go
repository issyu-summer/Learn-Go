package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制最大上传尺寸 8*1^20 = 8 mb
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "upload img failed")
		}
		c.SaveUploadedFile(file, file.Filename+"test.png")
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()

}

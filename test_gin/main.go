package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.POST("/upload", updateHandler)
	r.Run(":8000")
}

func updateHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}
func updateMultiHandler(c *gin.Context) {
	form, err := c.MultipartMemory()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
}

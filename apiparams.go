package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		// action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.Run(":8000")
}

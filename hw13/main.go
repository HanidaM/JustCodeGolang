package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/static/:file", func(c *gin.Context) {
		file := c.Param("file")
		c.File("hw13/files/" + file)
	})

	r.Run(":8080")
}

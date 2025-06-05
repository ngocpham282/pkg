package main

import (
	"pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.GET("/hek", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run(":8080")
}

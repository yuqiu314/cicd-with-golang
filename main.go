package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetMainEngine is an extract router for easy unit testing
func GetMainEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("London bridge is falling down"))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
func main() {
	GetMainEngine().Run(":80")
}

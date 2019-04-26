package main

import (
	"fmt"
	"os"

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
	// Heroku will have a PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	GetMainEngine().Run(":"+port)
}

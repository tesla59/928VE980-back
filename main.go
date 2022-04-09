package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Load default configs
	router = gin.Default()

	// Setup routers
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Server the App
	router.Run()
}

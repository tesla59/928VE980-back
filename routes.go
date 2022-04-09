package main

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes() {
	router.GET("/", Home)
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

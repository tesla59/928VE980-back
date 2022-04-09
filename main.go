package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Load default configs
	router = gin.Default()

	// Setup routers
	LoadRoutes()

	// Server the App
	router.Run()
}

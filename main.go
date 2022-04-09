package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Load default configs
	router = gin.Default()

	// Load all HTMLs
	router.LoadHTMLGlob("front/*.html")

	// Load all images
	LoadImages()

	// Setup routers
	LoadRoutes()

	// Server the App
	router.Run()
}

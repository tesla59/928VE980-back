package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var router *gin.Engine
var DB *gorm.DB
var Err error
var Confessions []Confession
var Count int64

func main() {

	// Load default configs
	router = gin.Default()

	// Load all HTMLs
	router.LoadHTMLGlob("front/*.html")

	// Load all images
	LoadImages()

	// Connect to DB
	ConnectDB()

	// Setup routers
	LoadRoutes()

	// Server the App
	router.Run()
}

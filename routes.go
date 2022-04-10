package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"net/http"
)

func LoadRoutes() {
	router.GET("/", Home)
	router.GET("/confessions", ConfessionHome)
	router.GET("/confessions/:id", SubmitConfession)
	router.GET("/team", Teams)
	router.POST("/thankyou/:id", POSTConfession)
}

func Home(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func ConfessionHome(c *gin.Context) {
	ID := xid.New()
	c.HTML(
		http.StatusOK,
		"confess.html",
		gin.H{
			"title": "Confession Page",
			"ID":    ID.String(),
		},
	)
}

func SubmitConfession(c *gin.Context) {
	ID := c.Param("id")
	c.HTML(
		http.StatusOK,
		"submit.html",
		gin.H{
			"title": "Confession page",
			"ID":    ID,
		},
	)
}

func Teams(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"team.html",
		gin.H{
			"title": "Our Team",
		},
	)
}

func POSTConfession(c *gin.Context) {
	// Migrate the schema
	DB.AutoMigrate(&Confession{})

	cnfs := Confession{
		ID:       c.Param("id"),
		To:       c.PostForm("to"),
		By:       c.PostForm("name"),
		Message:  c.PostForm("message"),
		SenderIP: c.ClientIP(),
		Posted:   false,
	}

	DB.Create(&cnfs)

	c.HTML(
		http.StatusOK,
		"thankyou.html",
		gin.H{
			"title": "Thank You",
		},
	)
}

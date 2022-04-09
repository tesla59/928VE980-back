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
	c.HTML(
		http.StatusOK,
		"submit.html",
		gin.H{
			"title": "Confession page",
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

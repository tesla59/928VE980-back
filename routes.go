package main

import (
	"encoding/json"
	"io"
	"log"

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
	router.GET("/memes", Memes)
	router.GET("/show", DisplayConfession)
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

func Memes(c *gin.Context) {
	// make GET request
	response, Err := http.Get("https://meme-api.herokuapp.com/gimme")
	if Err != nil {
		log.Println(Err)
	}

	body, _ := io.ReadAll(response.Body)
	response.Body.Close()

	var redd Reddit
	Err = json.Unmarshal(body, &redd)

	c.HTML(
		http.StatusOK,
		"meme_display.html",
		gin.H{
			"title": "Dank Maymay",
			"meme":  redd.URL,
		},
	)
}

func DisplayConfession(c *gin.Context) {
	DB.Model(&Confession{}).Count(&Count)
	DB.Limit(6).Offset(int(Count) - 6).Find(&Confessions)
	c.HTML(
		http.StatusOK,
		"confession_display.html",
		gin.H{
			"title": "Confessions",
			"f1":    Confessions[5].To,
			"b1":    Confessions[5].By,
			"m1":    Confessions[5].Message,
			"f2":    Confessions[4].To,
			"b2":    Confessions[4].By,
			"m2":    Confessions[4].Message,
			"f3":    Confessions[3].To,
			"b3":    Confessions[3].By,
			"m3":    Confessions[3].Message,
			"f4":    Confessions[2].To,
			"b4":    Confessions[2].By,
			"m4":    Confessions[2].Message,
			"f5":    Confessions[1].To,
			"b5":    Confessions[1].By,
			"m5":    Confessions[1].Message,
			"f6":    Confessions[0].To,
			"b6":    Confessions[0].By,
			"m6":    Confessions[0].Message,
		},
	)
}

package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadImages() {
	router.StaticFile("/anishL.jpeg", "./front/images/anishL.jpeg")
	router.StaticFile("/ayush.jpeg", "./front/images/ayush.jpeg")
	router.StaticFile("/divyansh.jpeg", "./front/images/divyansh.jpeg")
	router.StaticFile("/tesla.jpeg", "./front/images/tesla.jpeg")
	router.StaticFile("/conf.jpeg", "./front/images/CONFESSION.jpeg")
	router.StaticFile("/meme.jpeg", "./front/images/meme.jpeg")
}

func ConnectDB() error {
	var db *gorm.DB
	db, Err = gorm.Open(sqlite.Open("confession.db"), &gorm.Config{})
	if Err != nil {
		return Err
	}
	DB = db
	return nil
}

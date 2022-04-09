package main

func LoadImages() {
	router.StaticFile("/anishL.jpeg", "./front/images/anishL.jpeg")
	router.StaticFile("/ayush.jpeg", "./front/images/ayush.jpeg")
	router.StaticFile("/divyansh.jpeg", "./front/images/divyansh.jpeg")
	router.StaticFile("/tesla.jpeg", "./front/images/tesla.jpeg")
}
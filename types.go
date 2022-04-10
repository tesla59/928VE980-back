package main

type Confession struct {
	ID       string `gorm:"primaryKey"`
	By       string
	To       string
	Message  string
	SenderIP string
	Posted   bool
}

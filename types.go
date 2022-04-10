package main

type Confession struct {
	ID       string `gorm:"primaryKey"`
	By       string
	To       string
	Message  string
	SenderIP string
	Posted   bool
}

type Reddit struct {
	PostLink  string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
	Nsfw      bool     `json:"nsfw"`
	Spoiler   bool     `json:"spoiler"`
	Author    string   `json:"author"`
	Ups       int      `json:"ups"`
	Preview   []string `json:"preview"`
}

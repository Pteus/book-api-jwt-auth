package models

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Username string `json:"username"`
}

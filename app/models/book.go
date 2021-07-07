package models

type Book struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Abstract string `json:"abstract"`
}

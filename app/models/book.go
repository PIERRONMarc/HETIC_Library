package models

type BookRequest struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Abstract string `json:"abstract" binding:"required"`
}

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Abstract string `json:"abstract"`
}

func HydrateBookFromRequest(input BookRequest) Book {
	return Book{
		Title: input.Title,
		Author: input.Author,
		Abstract: input.Abstract,
	}
}

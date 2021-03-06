package controllers

import (
	"encoding/json"
	"hetic-library/models"
	"hetic-library/repositories"
	"hetic-library/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// search all books by title, author or abstract : GET /book/search
// queryParam query string
func SearchBooks(c *gin.Context) {
	query := c.Query("query")
	httpResponse, err := repositories.FindBooks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	var books models.Books
	searchResults, err := services.HttpResponseToSearchResults(httpResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	for _, result := range searchResults.Hits.Results {
		books = append(books, models.Book{
			ID:       result.ID,
			Title:    result.Data["title"].(string),
			Author:   result.Data["author"].(string),
			Abstract: result.Data["abstract"].(string),
		})
	}

	// endpoint response
	jsonResponse, err := json.Marshal(books)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.Data(http.StatusOK, "application/json", jsonResponse)
}

// Create and store a book : POST /book
func CreateBook(c *gin.Context) {

	var input models.BookRequest

	// Validation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusMethodNotAllowed, "Invalid input")
		return
	}

	// Store in Elasticsearch
	book := models.HydrateBookFromRequest(input)
	httpResponse, err := repositories.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	// Read body from Elasticsearch response to get document ID
	httpJsonResponse, err := services.HttpResponseToJson(httpResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	book.ID = httpJsonResponse["_id"].(string)

	// endpoint response
	jsonResponse, err := json.Marshal(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.Data(http.StatusCreated, "application/json", jsonResponse)
}

// Update a book : PUT /book/:book_id
func UpdateBook(c *gin.Context) {
	var input models.BookRequest

	// Validation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusMethodNotAllowed, "Invalid input")
		return
	}

	// Update in Elasticsearch
	book := models.HydrateBookFromRequest(input)
	httpResponse, err := repositories.UpdateBook(book, c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	} else if httpResponse.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusNotFound, "Book not found")
		return
	}

	book.ID = c.Param("book_id")

	// endpoint response
	jsonResponse, err := json.Marshal(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.Data(http.StatusOK, "application/json", jsonResponse)
}

// Delete a book : DELETE /book/:book_id
func DeleteBook(c *gin.Context) {
	httpResponse, err := repositories.DeleteBook(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	} else if httpResponse.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusNotFound, "Book not found")
		return
	}

	c.JSON(http.StatusOK, "Book deleted")
}

// Delete all books: DELETE /deleteAll
func DeleteAllBooks(c *gin.Context) {

	_, err := repositories.DeleteAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "All books have been deleted")
}

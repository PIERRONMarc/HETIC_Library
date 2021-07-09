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
func SearchBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello worlds",
	})
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
		c.JSON(http.StatusNotFound, "Document not found")
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

// Delete a book : PUT /book/:book_id
func DeleteBook(c *gin.Context) {
	var input models.BookRequest

	// Delete in Elasticsearch
	book := models.HydrateBookFromRequest(input)
	httpResponse, err := repositories.DeleteBook(book, c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
    } else if httpResponse.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusNotFound, "Document not found")
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

func DeleteAllBooks(c *gin.Context) {

	var input models.BookRequest

	// DeleteAll in Elasticsearch
	book := models.HydrateBookFromRequest(input)
	httpResponse, err := repositories.DeleteAllBooks(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
    } else if httpResponse.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusNotFound, "Document not found")
		return
	}
	// endpoint response
	jsonResponse, err := json.Marshal(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
    }

	c.Data(http.StatusOK, "application/json", jsonResponse)
}
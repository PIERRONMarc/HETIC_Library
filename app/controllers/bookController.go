package controllers

import (
	"github.com/gin-gonic/gin"
)

// search all books by title, author or abstract : GET /book/search
func SearchBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

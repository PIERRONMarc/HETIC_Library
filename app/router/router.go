package router

import (
	"hetic-library/controllers"
	"github.com/gin-gonic/gin"
)

func bookRoutes(r *gin.Engine) {
	r.GET("/book/search", controllers.SearchBooks)
	r.POST("/book", controllers.CreateBook)
	r.PUT("/book/:book_id", controllers.UpdateBook)
	r.DELETE("/book/:book_id", controllers.DeleteBook)
	r.DELETE("/deleteAll", controllers.DeleteAllBooks)
}

func SetRouter(r *gin.Engine) {
	bookRoutes(r)
}

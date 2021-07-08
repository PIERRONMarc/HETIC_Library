package router

import (
	"hetic-library/controllers"

	"github.com/gin-gonic/gin"
)

func bookRoutes(r *gin.Engine) {
	r.GET("/book/search", controllers.SearchBooks)
	r.POST("/book", controllers.CreateBook)
}

func SetRouter(r *gin.Engine) {
	bookRoutes(r)
}

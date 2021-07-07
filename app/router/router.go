package router

import (
	"github.com/gin-gonic/gin"
	"hetic-library/controllers"
)

func bookRoutes(r *gin.Engine) {
	r.GET("/book/search", controllers.SearchBooks)
}

func SetRouter(r *gin.Engine) {
	bookRoutes(r)
}

package main

import (
	"hetic-library/router"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()

	// TODO: insert fake data in elasticsearch

	router.SetRouter(r)

	r.Run(":8080")

	return r
}

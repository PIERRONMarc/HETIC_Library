package main

import (
	"hetic-library/router"
	"hetic-library/services/elasticsearch"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()
	elasticsearch.LoadFakeData()

	router.SetRouter(r)

	r.Run(":8080")

	return r
}

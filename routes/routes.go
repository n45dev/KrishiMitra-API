package routes

import (
	"github.com/gin-gonic/gin"
	"krishimitra-api/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, "pong")
	})

	router.GET("/crops", handlers.GetCrops)
	router.POST("/crops", handlers.PostCrops)

	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.PostProducts)

	router.GET("/news", handlers.GetNews)

	return router
}

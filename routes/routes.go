package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"krishimitra-api/docs"
	"krishimitra-api/handlers"
)

const version = "v0.1"
const basePath = "/api/" + version

func SetupRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = basePath

	router.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET(basePath+"/ping", func(c *gin.Context) {
		c.IndentedJSON(200, "pong")
	})

	router.GET(basePath+"/crops", handlers.GetCrops)
	router.POST(basePath+"/crops", handlers.PostCrops)

	router.GET(basePath+"/products", handlers.GetProducts)
	router.POST(basePath+"/products", handlers.PostProducts)

	router.GET(basePath+"/news", handlers.GetNews)

	return router
}

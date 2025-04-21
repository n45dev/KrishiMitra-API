package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"krishimitra-api/controllers"
	"krishimitra-api/docs"
	"krishimitra-api/middlewares"
)

const version = "v0.1"
const basePath = "/api/" + version

func SetupRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = basePath

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, "pong")
	})

	router.POST("/register", controllers.RegisterUser)

	router.POST("/login", controllers.LoginUser)

	p := router.Group(basePath)
	p.Use(middlewares.JwtAuthMiddleware())

	p.GET("/users/me", controllers.CurrentUser)

	p.GET("/crops", controllers.GetCrops)
	p.POST("/crops", controllers.PostCrops)

	p.GET("/products", controllers.GetProducts)
	p.POST("/products", controllers.PostProducts)

	p.GET("/news", controllers.GetNews)

	return router
}

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
	r := gin.Default()
	p := r.Group(basePath)
	p.Use(middlewares.JwtAuthMiddleware())

	docs.SwaggerInfo.BasePath = basePath

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, "pong")
	})

	r.POST("/register", controllers.RegisterUser)

	r.POST("/login", controllers.LoginUser)

	p.GET("/users/me", controllers.CurrentUser)

	p.GET("/crops", controllers.GetCrops)
	p.POST("/crops", controllers.PostCrops)

	p.GET("/products", controllers.GetProducts)
	p.POST("/products", controllers.PostProducts)

	p.GET("/news", controllers.GetNews)

	return r
}

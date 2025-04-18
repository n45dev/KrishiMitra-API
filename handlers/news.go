package handlers

import (
	"github.com/gin-gonic/gin"
	"krishimitra-api/db"
	"krishimitra-api/models"
	"net/http"
)

func GetNews(c *gin.Context) {
	var news []models.NewsArticle
	db.DB.Find(&news)
	c.IndentedJSON(http.StatusOK, news)
}

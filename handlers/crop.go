package handlers

import (
	"github.com/gin-gonic/gin"
	"krishimitra-api/db"
	"krishimitra-api/models"
	"net/http"
)

func GetCrops(c *gin.Context) {
	var crops []models.Crop
	db.DB.Find(&crops)
	c.IndentedJSON(http.StatusOK, crops)
}

func PostCrops(c *gin.Context) {
	var crop models.Crop
	if err := c.ShouldBindJSON(&crop); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&crop)
	c.IndentedJSON(http.StatusCreated, crop)
}

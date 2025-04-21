package controllers

import (
	"github.com/gin-gonic/gin"
	"krishimitra-api/db"
	"krishimitra-api/models"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)
	c.IndentedJSON(http.StatusOK, products)
}

func PostProducts(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&product)
	c.IndentedJSON(http.StatusCreated, product)
}

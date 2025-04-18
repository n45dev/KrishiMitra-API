package handlers

import (
	"github.com/gin-gonic/gin"
	"krishimitra-api/db"
	"krishimitra-api/models"
	"net/http"
)

// @BasePath /api/v1

// GetCrops godoc
// @Summary get crops
// @Schemes Crop
// @Description get all crops
// @Tags crops
// @Produce json
// @Success 200 {list} models.Crop
// @Router /crops [get]
func GetCrops(c *gin.Context) {
	var crops []models.Crop
	db.DB.Find(&crops)
	c.IndentedJSON(http.StatusOK, crops)
}

// @BasePath /api/v1

// PostCrops godoc
// @Summary post crops
// @Schemes Crop
// @Description post a crop
// @Tags crops
// @Accept json
// @Produce json
// @Success 200
// @Router /crops [post]
func PostCrops(c *gin.Context) {
	var crop models.Crop
	if err := c.ShouldBindJSON(&crop); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&crop)
	c.IndentedJSON(http.StatusCreated, crop)
}

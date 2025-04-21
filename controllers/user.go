package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"krishimitra-api/db"
	"krishimitra-api/models"
	"krishimitra-api/utils"
	"net/http"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(c *gin.Context) {
	var user models.UserRegister
	var userNew models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	userNew.Name = user.Name
	userNew.Email = user.Email
	userNew.HashedPassword = hashedPassword
	db.DB.Create(&userNew)
	c.IndentedJSON(http.StatusCreated, gin.H{"name": user.Name, "email": user.Email})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	var user LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u := models.User{}
	db.DB.Where("email = ?", user.Email).First(&u)

	if verifyPassword(user.Password, u.HashedPassword) {
		token, _ := utils.GenerateToken(u.ID)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func CurrentUser(c *gin.Context) {
	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	user := models.User{}
	db.DB.Where("id = ?", userId).First(&user)

	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

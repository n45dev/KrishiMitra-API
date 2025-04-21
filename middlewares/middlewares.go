package middlewares

import (
	"github.com/gin-gonic/gin"
	"krishimitra-api/utils"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

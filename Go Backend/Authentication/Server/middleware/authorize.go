package middleware

import (
	"auth-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("role", claims.Role)
		c.Next()
	}
}

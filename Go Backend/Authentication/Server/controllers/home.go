package controllers

import (
	"auth-go/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	claims, err := utils.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	fmt.Println(claims.RegisteredClaims)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the home page!", "role": claims.Role})
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Welcome to the home page",
		"message": "You have successfully logged in!",
	})
}

package routes

import (
	"auth-go/controllers"
	"github.com/gin-gonic/gin"
)

var WebsiteRoutes = func(c *gin.Engine) {
	c.POST("/login", controllers.Login)
	c.POST("/signup", controllers.SignUp)
	c.GET("/home", controllers.Home)
	c.POST("/reset", controllers.ForgetPassword)
}

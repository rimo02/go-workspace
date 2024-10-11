package main

import (
	"auth-go/database"
	"auth-go/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	r.LoadHTMLGlob("../Client/public/*")
	routes.WebsiteRoutes(r)
	err := r.Run(":9080")
	if err != nil {
		fmt.Println("Error running at port: 9080", err.Error())
	}

}

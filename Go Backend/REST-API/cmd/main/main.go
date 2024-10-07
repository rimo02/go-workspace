package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rimo02/rest-api/pkg/routes"
	"github.com/rimo02/rest-api/database"
)

func main() {
	router := gin.Default() // This creates an instance of gin.Engine
	routes.RegisterBookStoreRoutes(router)
    database.ConnectDB()
	router.Run(":9080")
}

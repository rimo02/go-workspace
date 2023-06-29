package main

import (
	"crm-fiber/database"
	"crm-fiber/lead"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:rimo398@/world?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database connected")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()

}

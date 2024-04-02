package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/putraprdn/go-fiber-crm/database"
	"github.com/putraprdn/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.NewLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(8000)
	defer database.DBConn.Close()
}

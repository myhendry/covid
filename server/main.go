package main

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to database
	database.Connect()

	app := fiber.New()

	// Use Middleware
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	
	// Setup Routes
	routes.Setup(app)
	
	app.Listen(":8000")

}
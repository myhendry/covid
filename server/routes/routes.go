package routes

import (
	"crypto.hendrylim/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/test/all", controllers.AllTests)
	app.Get("/api/test/:id", controllers.GetTest)
	app.Post("/api/test/add", controllers.AddTest)
	app.Put("/api/test/:id", controllers.UpdateTest)
	app.Delete("/api/test/:id", controllers.DeleteTest)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Yeah")
	})
}
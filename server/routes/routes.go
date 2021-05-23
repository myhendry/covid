package routes

import (
	"crypto.hendrylim/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/users", controllers.GetUsers)
	app.Post("/api/users/add", controllers.AddUser)

	app.Get("/api/tests", controllers.GetTests)
	app.Get("/api/tests/:id", controllers.GetTest)
	app.Post("/api/tests/add", controllers.AddTest)
	app.Put("/api/tests/:id", controllers.UpdateTest)
	app.Delete("/api/tests/:id", controllers.DeleteTest)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Yeah")
	})
}
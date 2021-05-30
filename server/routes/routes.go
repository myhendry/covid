package routes

import (
	"crypto.hendrylim/controllers"
	"crypto.hendrylim/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	tests := api.Group("/tests", middlewares.IsAuthenticated)

	tests.Get("/", controllers.GetTests)
	tests.Get("/try", controllers.TestTest)
	tests.Get("/:id", controllers.GetTest)
	tests.Post("/add", controllers.AddTest)
	tests.Post("/add2", controllers.AddTest2)
	tests.Put("/:id", controllers.UpdateTest)
	tests.Delete("/:id", controllers.DeleteTest)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Yeah")
	})
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)

	//* Authenticated Routes Below
	app.Use(middlewares.IsAuthenticated)

	app.Post("/api/auth/logout", controllers.Logout)
	app.Get("/api/auth/me", controllers.Me)
	app.Put("/api/auth/update-user", controllers.UpdateUser)
	app.Put("/api/auth/update-password", controllers.UpdatePassword)

	app.Delete("/api/all", controllers.DeleteAllTables)

	app.Get("/api/users", controllers.GetUsers)
	app.Post("/api/users/add", controllers.AddUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
	app.Delete("/api/users", controllers.DeleteAllUsers)

	app.Get("/api/roles", controllers.GetRoles)
	app.Post("/api/roles/add", controllers.AddRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)
	app.Delete("/api/roles", controllers.DeleteAllRoles)

	app.Get("/api/profiles", controllers.GetProfiles)
	app.Post("/api/profiles/add", controllers.AddProfile)
	app.Delete("/api/profiles/:id", controllers.DeleteProfile)
	app.Delete("/api/profiles", controllers.DeleteAllProfiles)

	app.Get("/api/books", controllers.GetBooks)
	app.Get("/api/books/:id", controllers.GetBook)
	app.Post("/api/books/add", controllers.AddBook)
	app.Put("/api/books/:id", controllers.UpdateBook)
	app.Delete("/api/books/:id", controllers.DeleteBook)
}
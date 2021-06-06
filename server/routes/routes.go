package routes

import (
	"crypto.hendrylim/controllers"
	"crypto.hendrylim/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	//! Order Routes
	orders := api.Group("/orders")

	orders.Get("/", controllers.GetOrders)

	ordersAuthenticated := orders.Use(middlewares.IsAuthenticated)

	ordersAuthenticated.Post("/add", controllers.AddOrder)
	ordersAuthenticated.Delete("/:id", controllers.DeleteOrder)
	ordersAuthenticated.Put("/:id", controllers.UpdateOrder)
	ordersAuthenticated.Post("/order-item/add", controllers.AddOrderItem)
	
	//! Test Routes
	tests := api.Group("/tests")

	tests.Get("/try", controllers.TestTest)

	testsAuthenticated := tests.Use(middlewares.IsAuthenticated)

	testsAuthenticated.Get("/", controllers.GetTests)
	testsAuthenticated.Get("/:id", controllers.GetTest)
	testsAuthenticated.Post("/add", controllers.AddTest)
	testsAuthenticated.Post("/add2", controllers.AddTest2)
	testsAuthenticated.Put("/:id", controllers.UpdateTest)
	testsAuthenticated.Put("/update/:id", controllers.UpdateTest2)
	testsAuthenticated.Delete("/:id", controllers.DeleteTest)
	//! Test Routes End

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Yeah")
	})
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)

	app.Post("/api/roles/add", controllers.AddRole)

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
	app.Get("/api/authors", controllers.GetAuthors)
	app.Post("/api/authors/add", controllers.AddAuthor)
}
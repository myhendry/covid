package controllers

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	database.DB.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	database.DB.Find(&book, id)

	return c.JSON(book)

}

func DeleteBook(c *fiber.Ctx) error {
		id := c.Params("id")

		var book models.Book

		database.DB.First(&book, id)

		if book.Title == "" {
			return c.Status(500).SendString("No Book Found with Given Id")
		}

		database.DB.Delete(&book)
		return c.SendString("Book Successfully Deleted")
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	database.DB.First(&book, id)

	if book.Title == "" {
		return c.Status(500).SendString("No Book Found with Given Id")
	}

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	database.DB.Model(&book).Updates(book)

	return c.SendString("Book Successfully Updated")
}

func AddBook(c *fiber.Ctx) error {
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	database.DB.Create(&book)

	return c.JSON(book)
}
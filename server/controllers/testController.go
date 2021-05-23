package controllers

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
)

func AllTests(c *fiber.Ctx) error {
	var tests []models.Test

	database.DB.Find(&tests)

	return c.JSON(tests)
}

func GetTest(c *fiber.Ctx) error {
	id := c.Params("id")

	var test models.Test

	database.DB.Find(&test, id)

	return c.JSON(test)

}

func DeleteTest(c *fiber.Ctx) error {
		id := c.Params("id")

		var test models.Test

		database.DB.First(&test, id)

		if test.Name == "" {
			return c.Status(500).SendString("No Test Found with Given Id")
		}

		database.DB.Delete(&test)
		return c.SendString("Test Successfully Deleted")

}

func UpdateTest(c *fiber.Ctx) error {
	id := c.Params("id")

	var test models.Test

	database.DB.First(&test, id)

	if test.Name == "" {
		return c.Status(500).SendString("No Test Found with Given Id")
	}

	if err := c.BodyParser(&test); err != nil {
		return err
	}

	database.DB.Model(&test).Updates(test)

	return c.SendString("Test Successfully Updated")
}

func AddTest(c *fiber.Ctx) error {
	var test models.Test

	if err := c.BodyParser(&test); err != nil {
		return err
	}

	database.DB.Create(&test)

	return c.JSON(test)
}
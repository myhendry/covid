package controllers

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func AddUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
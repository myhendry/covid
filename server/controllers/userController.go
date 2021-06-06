package controllers

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Role").Preload("Profile").Preload("Books.Authors").Find(&users)

	for i, user := range users {
		users[i].Nickname = user.GetNickname()
	}

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

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	database.DB.First(&user, id)

	if user.Name == "" {
		return c.Status(500).SendString("No User Found with Given Id")
	}

	database.DB.Delete(&user)
	return c.SendString("User Successfully Deleted")

}

func DeleteAllUsers(c *fiber.Ctx) error {
	database.DB.Exec("DELETE FROM users")
	return c.SendString("All Users Successfully Deleted")
}

func DeleteAllTables(c *fiber.Ctx) error {
	database.DB.Exec("DELETE FROM profiles")
	database.DB.Exec("DELETE FROM roles")
	database.DB.Exec("DELETE FROM users")
	return c.SendString("All Tables Successfully Deleted")
}
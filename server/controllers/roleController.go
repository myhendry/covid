package controllers

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func AddRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role

	database.DB.First(&role, id)

	if role.Name == "" {
		return c.Status(500).SendString("No Role Found with Given Id")
	}

	database.DB.Delete(&role)
	return c.SendString("Role Successfully Deleted")

}

func DeleteAllRoles(c *fiber.Ctx) error {
	database.DB.Exec("DELETE FROM roles")
	return c.SendString("All Roles Successfully Deleted")
}

// func DeleteUser(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	var user models.User

// 	database.DB.First(&user, id)

// 	if user.Name == "" {
// 		return c.Status(500).SendString("No User Found with Given Id")
// 	}

// 	database.DB.Delete(&user)
// 	return c.SendString("User Successfully Deleted")

// }

package controllers

import (
	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
)

func GetProfiles(c *fiber.Ctx) error {
	var profiles []models.Profile

	database.DB.Find(&profiles)

	return c.JSON(profiles)
}

func AddProfile(c *fiber.Ctx) error {
	var profile models.Profile

	if err := c.BodyParser(&profile); err != nil {
		return err
	}

	database.DB.Create(&profile)

	return c.JSON(profile)
}

func DeleteProfile(c *fiber.Ctx) error {
	id := c.Params("id")

	var profile models.Profile

	database.DB.First(&profile, id)

	if profile.Country == "" {
		return c.Status(500).SendString("No Profile Found with Given Id")
	}

	database.DB.Delete(&profile)
	return c.SendString("Profile Successfully Deleted")

}

func DeleteAllProfiles(c *fiber.Ctx) error {
	database.DB.Exec("DELETE FROM profiles")
	return c.SendString("All Profiles Successfully Deleted")
}
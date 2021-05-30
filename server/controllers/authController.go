package controllers

import (
	"errors"
	"strconv"
	"time"

	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"crypto.hendrylim/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	age, _ := strconv.Atoi(data["age"])
	roleId, _ := strconv.Atoi(data["role_id"])

	user := models.User{
		Name: data["name"],
		Age: uint(age),
		Email: data["email"],
		RoleID: uint(roleId),
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	errEmailNotFound := database.DB.Where("email = ?", data["email"]).First(&user).Error

	if (errors.Is(errEmailNotFound, gorm.ErrRecordNotFound)) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa((int(user.ID))))

	if err != nil {
		return c.Status(500).SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	 errUserNotFound := database.DB.Preload("Role").Where("id=?", id).First(&user).Error

	 if (errors.Is(errUserNotFound, gorm.ErrRecordNotFound)) {
		return c.Status(400).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// Cannot delete cookie. can only set expired time and empty the value to invalidate the cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,	
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged Out Successfully",
	})
}



package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func GetTests(c *fiber.Ctx) error {
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

		if test.Title == "" {
			return c.Status(500).SendString("No Test Found with Given Id")
		}

		database.DB.Delete(&test)
		return c.SendString("Test Successfully Deleted")

}

//* Using Structs Only
func UpdateTest(c *fiber.Ctx) error {
	id := c.Params("id")

	var test models.Test

	err := database.DB.First(&test, id).Error

	if (errors.Is(err, gorm.ErrRecordNotFound)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Test Not Found",
		})
	}

	if err := c.BodyParser(&test); err != nil {
		return err
	}

	database.DB.Model(&test).Updates(test)

	return c.SendString("Test Successfully Updated")
}

//* Using Structs Only
func AddTest(c *fiber.Ctx) error {
	var test models.Test

	if err := c.BodyParser(&test); err != nil {
		return err
	}

	database.DB.Create(&test)

	return c.JSON(test)
}

//* Using Maps and Empty Interfaces
func AddTest2(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	data["quantity"] = 1000

	res, err := createFromMap(data)

	if err != nil {
		fmt.Println(err)
	}

	database.DB.Create(&res)

	return c.JSON(res)
}

func createFromMap(m map[string]interface{}) (models.Test, error) {
	var result models.Test
	err := mapstructure.Decode(m, &result)
	return result, err
}

//* Using Where Clause in Find
func TestTest(c *fiber.Ctx) error {
	var tests []models.Test
	var count int64
	
	database.DB.Where("approved=?", true).Find(&tests).Count(&count)

	return c.JSON(fiber.Map{
		"tests": tests,
		"count": count,
	})
}

//* Using Maps and Empty Interfaces
func UpdateTest2(c *fiber.Ctx) error {
	id := c.Params("id")

	userId, _ := strconv.Atoi(id)
	
	var test map[string]interface{}

	if err := c.BodyParser(&test); err != nil {
		return err
	}

	s := models.Test{
		ID: uint(userId),
	}

	database.DB.Model(&s).Updates(test)

	return c.JSON(test)
	
}
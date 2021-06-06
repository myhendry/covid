package controllers

import (
	"errors"

	"crypto.hendrylim/database"
	"crypto.hendrylim/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Total = order.GetTotal()
	}

	return c.JSON(orders)
}

func AddOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return err
	}

	database.DB.Create(&order)

	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	err := database.DB.First(&order, id).Error

	if (errors.Is(err, gorm.ErrRecordNotFound)) {
				return c.Status(500).SendString("No Order Found with Given Id")
	}

	database.DB.Delete(&order)

	return c.JSON(fiber.Map{
		"message": "Order Successfully Deleted",
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	err := database.DB.First(&order, id).Error

	if (errors.Is(err, gorm.ErrRecordNotFound)) {
				return c.Status(500).SendString("No Order Found with Given Id")
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Parse Order",
		})
	}

	database.DB.Model(&order).Updates(order)

	return c.JSON(fiber.Map{
		"message": "Order Updated Successfully",
	})
}

func AddOrderItem(c * fiber.Ctx) error {
	var orderItem models.OrderItem

	if err := c.BodyParser(&orderItem); err != nil {
		return err
	}

	database.DB.Create(&orderItem)

	return c.JSON(fiber.Map{
		"message": "Order Item Added Successfully",
	})
}
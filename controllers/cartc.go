package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
)

func GetCartItems(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	var items []models.CartItem
	result := database.DB.Where("user_id = ?", userID).Find(&items)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

func AddToCart(c *fiber.Ctx) error {
	var item models.CartItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result := database.DB.Create(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

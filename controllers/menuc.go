package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"

	"github.com/gofiber/fiber/v2"
)

func GetMenuItems(c *fiber.Ctx) error {
	var items []models.MenuItem
	result := database.DB.Find(&items)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

func CreateMenuItem(c *fiber.Ctx) error {
	var item models.MenuItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result := database.DB.Create(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

func SearchMenuItemsByTags(c *fiber.Ctx) error {
	tag := c.Query("tag")
	if tag == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tag query parameter is required"})
	}

	var items []models.MenuItem
	result := database.DB.Where("tags LIKE ?", "%"+tag+"%").Find(&items)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(items)
}

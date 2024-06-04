package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
)

func GetSuggestions(c *fiber.Ctx) error {
	var suggestions []models.Suggestion
	result := database.Db.Find(&suggestions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(suggestions)
} // unpersonalised suggestions without the user id, means suggestions are not based on preferences but same for all

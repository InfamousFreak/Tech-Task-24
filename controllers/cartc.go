package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
)

func GetCartItems(c *fiber.Ctx) error {
	userID := c.Params("user_id")                                   //gets the user id parameter from the request url using c.params and stored in the variable userID
	var items []models.CartItem                                     //creates an empty slice to store the retrieved cart items
	result := database.DB.Where("user_id = ?", userID).Find(&items) //used to retrieve the cart items from the database matching the user id with the stored user id
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()}) //if error then 500 status code, JSON part creates a json response with a single key 'error' with error message for it
	}
	return c.Status(fiber.StatusOK).JSON(items) //if no error then status code 200, and JSON part sends the items slice as json response
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

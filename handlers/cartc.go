package handlers

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
	var item models.CartItem                    //variable item with type models.cartitem (presumably defines the structure of a cart item (e.g., it might include fields such as ProductID, Quantity, etc.).)
	if err := c.BodyParser(&item); err != nil { //BodyParser method of the Fiber context tries to read the JSON payload from the request body and unmarshal it into the item variable
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()}) //if error then sets status bad request and json response with single key "error" and err.Error() as its value
	}
	result := database.DB.Create(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item) // sends a json response of the items created
}

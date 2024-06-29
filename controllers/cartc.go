package controllers

import (
    "errors"
    "gorm.io/gorm"

	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
)
/*func GetCartItems(c *fiber.Ctx) error {
    // Get user ID from query parameter
    userID, err := c.ParamsInt("user_id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    // Check if the user exists
    var user models.UserProfile
    if err := database.Db.First(&user, userID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
    }

    // Retrieve cart items for the user
    var cartItems []models.CartItem
    if err := database.Db.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cart items"})
    }

    // If no items found, return an empty array instead of null
    if len(cartItems) == 0 {
        return c.JSON([]models.CartItem{})
    }

    return c.JSON(cartItems)
}*/

func GetCartItems(c *fiber.Ctx) error {
    // Get user ID from query parameter
    userID, err := c.ParamsInt("user_id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    // Check if the user exists
    var user models.UserProfile
    if err := database.Db.First(&user, userID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
    }

    // Retrieve cart items for the user, including menu item details
    var cartItemsWithDetails []struct {
        models.CartItem
        Name  string  `json:"item_name"`
        Price float64 `json:"item_price"`
    }

    if err := database.Db.Table("cart_items").
        Select("cart_items.*, menu_items.name, menu_items.price").
        Joins("LEFT JOIN menu_items ON cart_items.menu_item_id = menu_items.menu_item_id").
        Where("cart_items.user_id = ?", userID).
        Scan(&cartItemsWithDetails).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cart items"})
    }

    // If no items found, return an empty array instead of null
    if len(cartItemsWithDetails) == 0 {
        return c.JSON([]struct{}{})
    }

    return c.JSON(cartItemsWithDetails)
}

func AddToCart(c *fiber.Ctx) error {
    var cartItem models.CartItem
    if err := c.BodyParser(&cartItem); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Validate input
    if cartItem.UserID == 0 || cartItem.MenuItemID == 0 || cartItem.Quantity <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cart item data"})
    }

    // Check if the user exists
    var user models.UserProfile
    if err := database.Db.First(&user, cartItem.UserID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
    }

    // Check if the menu item exists
    var menuItem models.MenuItem
    if err := database.Db.First(&menuItem, "menu_item_id = ?", cartItem.MenuItemID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu item not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
    }

    // Check if the item is already in the cart
    var existingItem models.CartItem
    err := database.Db.Where("user_id = ? AND menu_item_id = ?", cartItem.UserID, cartItem.MenuItemID).First(&existingItem).Error
    if err == nil {
        // Item exists, update quantity
        existingItem.Quantity += cartItem.Quantity
        if err := database.Db.Save(&existingItem).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cart"})
        }
        cartItem = existingItem
    } else if errors.Is(err, gorm.ErrRecordNotFound) {
        // New item, create it
        if err := database.Db.Create(&cartItem).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to cart"})
        }
    } else {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
    }

    return c.Status(fiber.StatusCreated).JSON(cartItem)
}
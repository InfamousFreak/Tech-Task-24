package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func GetMenuItems(c *fiber.Ctx) error {
	var items []models.MenuItem
	result := database.Db.Find(&items) 
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
	result := database.Db.Create(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

/*func SearchMenuItemsByTags(c *fiber.Ctx) error {
	tag := c.Query("tag") // this retrieves the query parameter named "tag" from the HTTP request, if fails to retrieve, an empty string will be returned
	if tag == "" {        //checks if the string is empty
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tag query parameter is required"})
	}

	var items []models.MenuItem
	result := database.Db.Where("LOWER(tags) LIKE LOWER(?)", "%"+tag+"%").Find(&items)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(items) //if no error then sends a json response
}*/
func SearchMenuItemsByTags(c *fiber.Ctx) error {
    partialTag := c.Query("tag")
    if partialTag == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tag query parameter is required"})
    }

    var items []models.MenuItem
    // Use LOWER() for case-insensitive search and LIKE for partial matching
    result := database.Db.Where("LOWER(tags) LIKE LOWER(?)", "%"+partialTag+"%").Find(&items)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(items)
}


func UpdateMenuItem(c *fiber.Ctx) error {
	var updates models.MenuItem
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")

	// Find the menu item by its ID
	var menuItem models.MenuItem
	if err := database.Db.First(&menuItem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "MenuItem not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Update the fields in the database
	if err := database.Db.Model(&menuItem).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(menuItem)
}
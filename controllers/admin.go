package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/gofiber/fiber/v2"
)

//functiont o get all the registered user profiles
func ShowProfiles(c *fiber.Ctx) error {
    var users []models.UserProfile

    // Query the database for all user profiles
    result := database.Db.Find(&users)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
        "data":    users,
        "success": true,
        "message": "Retrieved Successfully",
    })
}

//function to get a specific user profile
func AdminGetUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var userProfile models.UserProfile
	result := database.Db.First(&userProfile, id)
	if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "UserProfile not found"})
    }
    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"data":    userProfile,
			"success": true,
			"message": "Retrieved Successfully",
	})
	return nil
}

//function for admin to update thr user profiel
func AdminUpdateUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var userProfile models.UserProfile
	result := database.Db.First(&userProfile, id)
	if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "UserProfile not found"})
    }
	if err := c.BodyParser(&userProfile); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    database.Db.Save(&userProfile)
    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": userProfile,
		"success": true,
		"message": "Updated Successfully",
	})
	return nil
}

func AdminDeleteUserProfile(c *fiber.Ctx) error {
    id := c.Params("id")

    // Find the user profile to ensure it exists
    var userProfile models.UserProfile
    if err := database.Db.First(&userProfile, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "UserProfile not found"})
    }

    // Delete the user profile
    if err := database.Db.Delete(&userProfile, id).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user profile"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
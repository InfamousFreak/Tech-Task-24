package controllers

import (
	"fmt"

	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
	"github.com/InfamousFreak/Tech-Task-24/passwordhashing"
)

func CreateUserProfile(c *fiber.Ctx) error {
	newUser := new(models.UserProfile)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

    if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "Error": "Username, email, and password are required",
        })
    }

	hashedPassword, err := passwordhashing.HashPassword(newUser.Password)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Failed to hash password"})
    }
    newUser.Password = hashedPassword

	createResult := database.Db.Create(&newUser)
	if createResult.Error != nil {
		c.Status(400).JSON(&fiber.Map{"error": createResult.Error.Error()})
		return createResult.Error
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    newUser,
		"success": true,
		"message": "Created Successfully",
	})
	return nil
}


/*func UpdateUserProfile(c *fiber.Ctx) error {
	var user models.UserProfile
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result := database.DB.Save(&user) //saving the value of given struct as record, updating an existing record in database or creating a new
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}*/

func UpdateUserProfile(c *fiber.Ctx) error {
    var newUser models.UserProfile
    userID := database.Convert(c.Params("id"))
    fmt.Println(userID)

    if err := c.BodyParser(&newUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
    }

    newUser.ID = userID // Update this line to use ID instead of UserID

    var existingUser models.UserProfile
    if err := database.Db.First(&existingUser, "id=?", userID).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{"Error": err.Error})
    }

    database.Db.First(&existingUser, userID)
    if err := database.Db.First(&existingUser, "id=?", userID).Error; err != nil { // Update this line to use ID instead of user_id
        result := database.Db.Create(&newUser)
        if result.Error != nil {
            c.Status(400).JSON(&fiber.Map{"error": result.Error.Error()})
        }
    } else {
        result := database.Db.Model(&existingUser).Updates(newUser)
        if result.Error != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": result.Error.Error()})
        }
    }

    return c.Status(200).JSON(&fiber.Map{
        "data":    newUser,
        "success": true,
        "message": "Updated Successfully",
    })
}

func DeleteUserProfile(c *fiber.Ctx) error {
    userID := database.Convert(c.Params("id"))

    var existingUser models.UserProfile
    if err := database.Db.First(&existingUser, "id=?", userID).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    result := database.Db.Delete(&existingUser)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
    }

    if result.RowsAffected == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    return c.Status(200).JSON(fiber.Map{
        "success": true,
        "message": "User deleted successfully",
    })
}

func SelectRole(c *fiber.Ctx) error {

    var roleData struct {
        UserID       uint   `json:"user_id"`
        Role         string `json:"role"`
        BusinessLicenseNumber string `json:"business_license_number,omitempty"`
        RestaurantType string `json:"restaurant_type,omitempty"`
    }

    if err := c.BodyParser(&roleData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    var userProfile models.UserProfile
    if err := database.Db.First(&userProfile, roleData.UserID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    userProfile.Role = roleData.Role

    if err := database.Db.Save(&userProfile).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    if roleData.Role == "restaurateur" {
        restaurateur := models.Admin{
            UserProfileID: userProfile.UserID,
            BusinessLicenseNumber: roleData.BusinessLicenseNumber,
            RestaurantType: roleData.RestaurantType,
        }

        if err := database.Db.Create(&restaurateur).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Role updated successfully"})
}




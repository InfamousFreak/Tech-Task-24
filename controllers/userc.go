package controllers

import (
	"fmt"

	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUserProfile(c *fiber.Ctx) error {
	newUser := new(models.UserProfile)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
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

// Check if the user already exists
//result := database.DB.First(&newUser, "username = ?", userName)
//if result.Error != nil {
//	if result.Error == gorm.ErrRecordNotFound {
// If not found, create a new user profile
/*createResult := database.DB.Create(&newUser)
			if createResult.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": createResult.Error.Error()})
			}
			return c.Status(fiber.StatusCreated).JSON(newUser)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// If user exists, return the user profile
	return c.Status(fiber.StatusOK).JSON(newUser)
}

/*func Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	result := database.Db.Create(&newUser)
	//database.Db.Create(&newUser)
	if result.Error != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": result.Error,
		})
		return result.Error
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    newUser,
		"success": true,
		"message": "Successfully registered",
	})
	return nil
}*/

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
	userID := database.Convert(c.Params("user_id"))
	fmt.Println(userID)
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	newUser.UserID = userID
	var existingUser models.UserProfile

	if err := database.Db.First(&existingUser, "user_id=?", userID).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": err.Error})
	}
	database.Db.First(&existingUser, userID)

	if err := database.Db.First(&existingUser, "user_id=?", userID).Error; err != nil {
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
    userID := database.Convert(c.Params("user_id"))

    var existingUser models.UserProfile
    if err := database.Db.First(&existingUser, "user_id=?", userID).Error; err != nil {
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


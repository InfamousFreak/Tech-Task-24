package controllers

import (
    "time"
    "errors"
    "gorm.io/gorm"

	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/gofiber/fiber/v2"
    "github.com/InfamousFreak/Tech-Task-24/config"
    jwt "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

func AdminSignup(c *fiber.Ctx) error {
    admin := new(models.Admin)
    if err := c.BodyParser(admin); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to hash password",
        })
    }
    admin.Password = string(hashedPassword)

    // Save the admin to the database
    if err := database.Db.Create(admin).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create admin account",
        })
    }

	claims := jwt.MapClaims{
		"ID": admin.ID,
		"email": admin.Email,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.AdminSecret))

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Failed to generate token",
            "error":   "Internal server error",
        })
    }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "message": "Admin Created successfully",
        "data": fiber.Map{
            "user": sanitizeAdminData(admin),
            "token": t,
        },
    })
}

func sanitizeAdminData(admin *models.Admin) fiber.Map {
    return fiber.Map{
        "id":              admin.ID,
        "name":            admin.Name,
        "email":           admin.Email,
    }
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

func GetAllAdminProfiles(c *fiber.Ctx) error {
    var admins []models.Admin

    if err := database.Db.Find(&admins).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch admin profiles",
        })
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
        "data":    admins,
        "success": true,
        "message": "Retrieved Successfully",
    })
}

func DeleteAdminProfile(c *fiber.Ctx) error {
    adminID := c.Params("id")

    var admin models.Admin
    if err := database.Db.First(&admin, adminID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Admin not found",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to find admin",
        })
    }

    if err := database.Db.Delete(&admin).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete admin",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Admin deleted successfully",
    })
}
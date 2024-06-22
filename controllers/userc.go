package controllers

import (
	"fmt"
    "time"

	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
	"github.com/InfamousFreak/Tech-Task-24/passwordhashing"
    "github.com/InfamousFreak/Tech-Task-24/config"
    jtoken "github.com/golang-jwt/jwt/v4"
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
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Failed to create user",
            "error":   "Database error",
        })
    }

    //return nil
    claims := jtoken.MapClaims{
        "ID":  newUser.UserID,
        "email": newUser.Email,
        "exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }
    token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
    t,err:= token.SignedString([]byte(config.Secret))

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Failed to generate token",
            "error":   "Internal server error",
        })
    }

    return c.JSON(models.LoginResponse{
		Token:t,
	})

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "message": "User created successfully",
        "data": fiber.Map{
            "user":  sanitizeUserData(newUser),
            "token": token,
        },
    })

}

func sanitizeUserData(user *models.UserProfile) *models.UserProfile {
    return &models.UserProfile{
        UserID:    user.UserID,
        Name:      user.Name,
        Email:     user.Email,
    }
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

/*func SelectRole(c *fiber.Ctx) error {

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
}*/

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

func Role(c *fiber.Ctx) error {
    var currUser models.UserProfile
    var currRole models.Roles

    if err := c.BodyParser(&currRole); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "Error": err.Error(),
        })
    }

    user := c.Locals("user").(*jtoken.Token)
    claims := user.Claims.(jtoken.MapClaims)
    idF6 := claims["ID"].(float64)
    id := uint(idF6)

    if result := database.Db.First(&currUser, id).Error; result != nil {
        return c.Status(400).JSON(&fiber.Map{
            "message": "User does not exist",
        })
    }

    currUser.Role = currRole.Role

    if currRole.Role == "restaurateur" {
        if currRole.BusinessLicense == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "message": "Business license required for restaurateurs",
            })
        }
        currUser.BusinessLicense = currRole.BusinessLicense
    }

    database.Db.Save(&currUser)

    day := time.Hour * 24
    claims = jtoken.MapClaims{
        "ID":    currUser.ID,
        "email": currUser.Email,
        "role":  currUser.Role,
        "expi":  time.Now().Add(day * 1).Unix(),
    }

    token2 := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
    t, err := token2.SignedString([]byte(config.Secret))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    /*response := models.LoginResponse{
        Token: t,
    }*/

    // Determine redirect URL based on role
    var redirectURL string
    if currUser.Role == "customer" {
        redirectURL = "/customer.html"
    } else if currUser.Role == "restaurateur" {
        redirectURL = "/restaurateur.html"
    }

    // Add redirection URL to the response
    return c.JSON(fiber.Map{
        "token": t,
        "redirect": redirectURL,
    })
}





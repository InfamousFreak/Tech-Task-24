package controllers

import (
	"fmt"
    "time"
    "errors"
    "gorm.io/gorm"

	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/gofiber/fiber/v2"
	"github.com/InfamousFreak/Tech-Task-24/passwordhashing"
    "github.com/InfamousFreak/Tech-Task-24/config"
    jwt "github.com/golang-jwt/jwt/v4"
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

    // Check if user already exists
    var existingUser models.UserProfile
    if err := database.Db.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "success": false,
            "message": "User with this email already exists",
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

    claims := jwt.MapClaims{
        "ID":    newUser.ID, // Use the default ID from gorm.Model
        "email": newUser.Email,
        "exp":   time.Now().Add(time.Hour * 12).Unix(), // Token expires in 12 hours
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(config.Secret))

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Failed to generate token",
            "error":   "Internal server error",
        })
    }

    c.Cookie(&fiber.Cookie{
        Name: "token",
        Value: t,
        Expires: time.Now().Add(time.Hour * 12),
        HTTPOnly: true,
        Secure: true,
        SameSite: "Strict",
    })

    // The following return statement will not be executed, consider removing it
    // return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    //     "success": true,
    //     "message": "User created successfully",
    //     "data": fiber.Map{
    //         "user":  sanitizeUserData(newUser),
    //         "token": t,
    //     },
    // })

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "message": "User Created successfully",
        "data": fiber.Map{
            "user": sanitizeUserData(newUser),
            "token": t,
        },
    })
}

/*func sanitizeUserData(user *models.UserProfile) *models.UserProfile {
    return &models.UserProfile{
        Model:  user.Model, // Include gorm.Model to get ID
        Name:   user.Name,
        Email:  user.Email,
        City: user.City,
        Role: user.Role,
        BusinessLicense: user.BusinessLicense,
    }
}*/
func sanitizeUserData(user *models.UserProfile) fiber.Map {
    return fiber.Map{
        "id":              user.ID,
        "name":            user.Name,
        "email":           user.Email,
        "city":            user.City,
        "role":            user.Role,
        "businessLicense": user.BusinessLicense,
    }
}

func ShowUserProfile(c *fiber.Ctx) error {
    userID := c.Params("id")
    if userID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success": false,
            "message": "User ID is required",
        })
    }

    var user models.UserProfile
    result := database.Db.First(&user, userID)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "success": false,
                "message": "User not found",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Error fetching user profile",
        })
    }

    // Sanitize user data before sending
    sanitizedUser := sanitizeUserData(&user)

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success": true,
        "message": "User profile fetched successfully",
        "data":    sanitizedUser,
    })
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

/*func Role(c *fiber.Ctx) error {
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
    }

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
}*/






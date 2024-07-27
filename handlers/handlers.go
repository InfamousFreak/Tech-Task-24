package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	//"errors"
	"gorm.io/gorm"
	
	//"golang.org/x/crypto/bcrypt"
	"github.com/InfamousFreak/Tech-Task-24/config"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/InfamousFreak/Tech-Task-24/repository"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/InfamousFreak/Tech-Task-24/database"
)

// Login route
func Login(c *fiber.Ctx) error {
	
	loginRequest := new(models.LoginRequest) 
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(), 
		})
	}
	// find the user by repository.find database query
    user, err := repository.Find(database.Db, loginRequest.Email, loginRequest.Password) 
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid Credentials", 
        })
    }

	tokenExpiration := time.Now().Add(time.Hour * 12)
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   tokenExpiration.Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {                                     
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	// return the JWT token and the userid for cart fetching in the response body
	return c.JSON(models.LoginResponse{
		Token: t,
		UserID: user.ID,
	})
}


// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return c.SendString("Welcome" + email)
}

func GoogleLogin(c *fiber.Ctx) error {
    url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
    return c.Redirect(url, fiber.StatusTemporaryRedirect)
}


func GoogleCallback(c *fiber.Ctx) error {
    state := c.Query("state")
    if state != "randomstate" {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid state")
    }

    code := c.Query("code")
    googleConfig := config.GoogleConfig()
    token, err := googleConfig.Exchange(context.Background(), code)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Code-Token exchange failed")
    }

    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user data")
    }
    defer resp.Body.Close()

    userData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to read user data")
    }

    var user models.UserProfile
    if err := json.Unmarshal(userData, &user); err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to parse user data")
    }

    // Check if user exists, if not create a new user
    var existingUser models.UserProfile
    result := database.Db.Where("email = ?", user.Email).First(&existingUser)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // Create new user
            if err := database.Db.Create(&user).Error; err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "success": false,
                    "message": "Failed to create user",
                    "error":   err.Error(),
                })
            }
        } else {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "success": false,
                "message": "Database error",
                "error":   result.Error.Error(),
            })
        }
    } else {
        user = existingUser
    }

    // Generate JWT
    jwtToken, err := generateJWT(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Failed to generate token",
            "error":   err.Error(),
        })
    }

    return c.JSON(models.LoginResponse{
        Token: jwtToken,
    })
}

func generateJWT(user models.UserProfile) (string, error) {
    day := time.Hour * 24
    claims := jwt.MapClaims{
        "ID":    user.ID,
        "email": user.Email,
        "exp":   time.Now().Add(day * 1).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.Secret))
}
func AdminLogin(c *fiber.Ctx) error {
    loginRequest := new(models.AdminLoginRequest)
    if err := c.BodyParser(loginRequest); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    // Find the admin using the new FindAdmin function
    admin, err := repository.FindAdmin(database.Db, loginRequest.Email, loginRequest.Password)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid credentials",
        })
    }

    // Generate JWT token
    tokenExpiration := time.Now().Add(time.Hour * 12)
    claims := jwt.MapClaims{
        "ID":    admin.ID,
        "email": admin.Email,
        "role":  "admin",
        "exp":   tokenExpiration.Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(config.Secret))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to generate token",
        })
    }
    // Return the JWT token
    return c.JSON(models.AdminLoginResponse{
        Token:   t,
        AdminID: admin.ID,
    })
}


func ProtectedAdmin(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    role := claims["role"].(string)
    
    if role != "admin" {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "Access denied",
        })
    }

    // Admin-specific logic here
    return c.SendString("Welcome, Admin")
}	



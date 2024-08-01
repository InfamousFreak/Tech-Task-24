package handlers

import (
	"context"
	"io"
	"net/http"
	"time"
	"encoding/json"
	//"errors"
    //                      "fmt"
    //"gorm.io/gorm"
	
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



/*func GoogleCallback(c *fiber.Ctx) error {
    state := c.Query("state")
    if state != "randomstate" {
        return c.SendString("States don't Match!!")
    }


    code := c.Query("code")
    googlecon := config.GoogleConfig()
    token, err := googlecon.Exchange(context.Background(), code)
   if err != nil {
   return c.SendString("Code-Token Exchange Failed")
    }
    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
   if err != nil {
   return c.SendString("Cannot retrieve the user data")
    }
   defer resp.Body.Close()
    userData, err := io.ReadAll(resp.Body)
   if err != nil {
   return c.SendString("JSON Parsing Failed")
    }
   var newUser models.UserProfile
    json.Unmarshal(userData, &newUser)
   if err := database.Db.Where("email = ?", newUser.Email).First(&newUser).Error; err != nil {
    result := database.Db.Create(&newUser)
   if result.Error != nil {
    c.Status(400).JSON(&fiber.Map{
   "data": nil,
   "success": false,
   "message": result.Error,
    })
   return result.Error
    }
    database.Db.Save(&newUser)
    }
    day := time.Hour * 24
    claims := jwt.MapClaims{
   "ID": newUser.ID,
   "email": newUser.Email,
   "expi": time.Now().Add(day * 1).Unix(),
    }
    token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token2.SignedString([]byte(config.Secret))
   if err != nil {
   return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    fmt.Println()
   return c.JSON(models.LoginResponse{
    Token: t,
    })
   }*/

    func GoogleCallback(c *fiber.Ctx) error {
        state := c.Query("state")
        if state != "randomstate" {
            return c.SendString("States don't Match!!")
        }

        code := c.Query("code")

        googlecon := config.GoogleConfig()

        token, err := googlecon.Exchange(context.Background(), code)
        if err != nil {
            return c.SendString("Code-Token Exchange Failed")
        }

        resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
        if err != nil {
            return c.SendString("Cannot retrieve the user data")
        }
        defer resp.Body.Close()

        userData, err := io.ReadAll(resp.Body)
        if err != nil {
            return c.SendString("JSON Parsing Failed")
        }
        var User models.UserProfile
        json.Unmarshal(userData, &User)
        if err := database.Db.Where("email = ?", User.Email).First(&User).Error; err != nil {
            result := database.Db.Create(&User)
            if result.Error != nil {
                c.Status(400).JSON(&fiber.Map{
                    "data":    nil,
                    "success": false,
                    "message": result.Error,
                })
                return result.Error                                             
            }
            database.Db.Save(&User)
        }

        day := time.Hour * 24
        claims := jwt.MapClaims{
            "ID":    User.ID,
            "email": User.Email,
            "name" : User.Name,
            "expi":  time.Now().Add(day * 1).Unix(),
        }
        token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        t, err := token2.SignedString([]byte(config.Secret))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        c.Cookie(&fiber.Cookie{
            Name:     "token",
            Value:    t,
            Expires:  time.Now().Add(day * 1),
        })

        redirectUrl := "http://127.0.0.1:5501/frontend/customer.html"
	    return c.Redirect(redirectUrl, fiber.StatusSeeOther)

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


/*func CheckAuth(c *fiber.Ctx) error {
    cookie := c.Cookies("token")
    if cookie == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "authenticated": false,
        })
    }

    // Verify the JWT token here
    // If valid, return success. If not, return unauthorized

    return c.JSON(fiber.Map{
        "authenticated": true,
    })
}*/


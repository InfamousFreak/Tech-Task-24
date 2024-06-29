package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	//"fmt"

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


	if user.Role == "restaurateur" && user.BusinessLicense != loginRequest.BusinessLicense {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid license number",
        })
    }

	tokenExpiration := time.Now().Add(time.Hour * 12)
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"role": user.Role,
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

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}
	 var newUser models.UserProfile
	 json.Unmarshal(userData,&newUser)
	 if err:=database.Db.Where("email = ?", newUser.Email).First(&newUser).Error;err!=nil{
		result := database.Db.Create(&newUser)
		if result.Error != nil {
			c.Status(400).JSON(&fiber.Map{
				"data":    nil,
				"success": false,
				"message": result.Error,
			})
			return result.Error
		}
		if result.Error != nil {
			c.Status(400).JSON(&fiber.Map{
				"data":    nil,
				"success": false,
				"message": result.Error,
			})
			return result.Error
		}
		//return nil,errors.New("user is not found")
		newUser.Role="customer";
	}
	

	day:=time.Hour*24;
	claims:=jwt.MapClaims{
		"ID": newUser.ID,
		"email":newUser.Email,
		"role":newUser.Role,
		"expi":time.Now().Add(day*1).Unix(),
	}
	token2:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	t,err:=token2.SignedString([]byte(config.Secret))
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error":err.Error(),})
	}
	return c.JSON(models.LoginResponse{
		Token:t,
	})

}



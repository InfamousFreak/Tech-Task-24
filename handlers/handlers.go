package handlers

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/InfamousFreak/Tech-Task-24/config"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/InfamousFreak/Tech-Task-24/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	//"github.com/InfamousFreak/Tech-Task-24/passwordhashing"
	"github.com/InfamousFreak/Tech-Task-24/database"
)

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest) //new is a go operation that allocates memory for the variable and returns a pointer to it
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(), //This part parses the request body into a LoginRequest struct. If parsing fails, it returns a 400 Bad Request status with the error message.
		})
	}
	// Find the user by credentials
	user, err := repository.Find(loginRequest.Email, loginRequest.Password) //interacts with a repository to find a user based on their email and password
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(), //This part calls a repository function Find to verify the user's email and password. If the credentials are incorrect, it returns a 401 Unauthorized status with the error message.
		})
	}
	//var user models.UserProfile
    if err := database.Db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
    }

    // Compare the provided password with the stored hashed password
    //if err := passwordhashing.CompareHashAndPassword(loginRequest.Password, user.Password); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
    	

	day := time.Hour * 24 
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"city":  user.City,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims) 
	
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {                                     
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the JWT token in the response body
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	city := claims["city"].(string)
	return c.SendString("Welcome 👋" + email + " " + city)
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
		return c.SendString("Code-Token Exchange failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("Code-Data fetch failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(userData))

}

package main

import (
	"github.com/InfamousFreak/Tech-Task-24/config"
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/handlers"
	"github.com/InfamousFreak/Tech-Task-24/middlewares"
	"github.com/InfamousFreak/Tech-Task-24/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	dbEr := database.InitDB()
	if dbEr != nil {
		panic("Failed connecting to the database")
	}
	//send a string back for GET CALLS to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	app.Use(cors.New(cors.Config{
        AllowCredentials: true,
        AllowOrigins: "http://127.0.0.1:5500", 
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))
	app.Static("/", "/frontend")

	jwt := middlewares.AuthMiddleware()
	config.GoogleConfig()
	//config.GithubConfig()

    app.Get("/dashboard", jwt, func(c *fiber.Ctx) error {
        userID := c.Locals("userID")
        return c.JSON(fiber.Map{
            "success": true,
            "message": "Welcome to your dashboard",
            "userID":  userID,
        })
    })
    
	app.Get("/google_login", handlers.GoogleLogin)
	app.Get("/google_callback", handlers.GoogleCallback)
	app.Post("/login", handlers.Login)
	app.Get("/protected", jwt, handlers.Protected)
	routes.SetupRouter(app)
	app.Listen(":8080")

}

/*func selectRole(c *fiber.Ctx) error {
    var roleSelection struct {
        UserID         uint   `json:"user_id"`
        Role           string `json:"role"`
        BusinessLicense string `json:"business_license"`
    }
    if err := c.BodyParser(&roleSelection); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    var user UserProfile
    if err := db.First(&user, roleSelection.UserID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "user not found"})
    }
    user.Role = roleSelection.Role
    if roleSelection.Role == "Restaurateur" {
        if roleSelection.BusinessLicense == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "business license required"})
        }
        user.BusinessLicense = roleSelection.BusinessLicense
    }
    db.Save(&user)
    if roleSelection.Role == "Customer" {
        return c.Redirect("/customer.html")
    } else if roleSelection.Role == "Restaurateur" {
        return c.Redirect("/restaurateur.html")
    }
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "invalid role"})
}*/


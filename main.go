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
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
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
    app.Post("/admin/login", handlers.AdminLogin)
    app.Post("/admin/protected", handlers.ProtectedAdmin)
	routes.SetupRouter(app)
	app.Listen(":8080")

}



package main

import (
	"github.com/InfamousFreak/Tech-Task-24/config"
	"github.com/InfamousFreak/Tech-Task-24/handlers"
	"github.com/InfamousFreak/Tech-Task-24/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	config.GoogleConfig()
	//config.GithubConfig()

	app.Get("/google_login", handlers.GoogleLogin)
	app.Post("/google_callback", handlers.GoogleCallback)
	//app := fiber.New()

	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	// Create a Login route
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)
	// Listen on port 8080
	app.Listen(":8080")

}

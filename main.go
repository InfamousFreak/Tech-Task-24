package main

import (
	"github.com/InfamousFreak/Tech-Task-24/config"
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/InfamousFreak/Tech-Task-24/handlers"
	"github.com/InfamousFreak/Tech-Task-24/middlewares"
	"github.com/InfamousFreak/Tech-Task-24/routes"
	"github.com/gofiber/fiber/v2"
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

	jwt := middlewares.NewAuthMiddleware(config.Secret)
	config.GoogleConfig()
	//config.GithubConfig()
	app.Get("/google_login", handlers.GoogleLogin)
	app.Get("/google_callback", handlers.GoogleCallback)
	app.Post("/login", handlers.Login)
	app.Get("/protected", jwt, handlers.Protected)
	routes.SetupRouter(app)
	app.Listen(":8080")

}

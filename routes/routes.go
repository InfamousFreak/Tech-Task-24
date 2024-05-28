package handlers

import (
	"github.com/InfamousFreak/Tech-Task-24/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/menu", handlers.GetMenuItems)
	app.Post("/menu", handlers.CreateMenuItem)
	app.Get("/menu/search", handlers.SearchMenuItemsByTags) //new search route for tags

	app.Get("/profile/:id", handlers.GetUserProfile)
	app.Get("/profile/:id", handlers.UpdateUserProfile)

	app.Get("/cart", handlers.AddToCart)
	app.Get("/cart/:user_id", handlers.GetCartItems)

	app.Get("/suggestions", handlers.GetSuggestions)

}

package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/menu", controllers.GetMenuItems)
	app.Post("/menu", controllers.CreateMenuItem)
	app.Get("/menu/search", controllers.SearchMenuItemsByTags) //new search route for tags

	app.Get("/profile/:id", controllers.GetUserProfile)
	app.Get("/profile/:id", controllers.UpdateUserProfile)

	app.Get("/cart", controllers.AddToCart)
	app.Get("/cart/:user_id", controllers.GetCartItems)

	app.Get("/suggestions", controllers.GetSuggestions)

}

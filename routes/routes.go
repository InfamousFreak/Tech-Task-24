package routes

import (
	"github.com/InfamousFreak/Tech-Task-24/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/menu/get", controllers.GetMenuItems)
	app.Post("/menu/create", controllers.CreateMenuItem)
	app.Get("/menu/search", controllers.SearchMenuItemsByTags) //new search route for tags
	app.Get("/menu/search/name", controllers.SearchMenuItemsByName) //new search route for searching food items by name

	app.Post("/profile/create", controllers.CreateUserProfile)
	app.Post("/profile/:user_id", controllers.UpdateUserProfile)
	app.Delete("/profile/:user_id", controllers.DeleteUserProfile)

	app.Get("/cart", controllers.AddToCart)
	app.Get("/cart/:user_id", controllers.GetCartItems)

	app.Get("/suggestions/:user_id", controllers.GetSuggestions)

}

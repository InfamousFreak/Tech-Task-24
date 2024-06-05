package routes

import (
	"github.com/InfamousFreak/Tech-Task-24/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/menu/get", controllers.GetMenuItems)
	app.Post("/menu/create", controllers.CreateMenuItem)
	app.Get("/menu/search", controllers.SearchMenuItemsByTags) //new search route for tags

	app.Post("/profile/create", controllers.CreateUserProfile)
	app.Post("/profile/:id", controllers.UpdateUserProfile)
	app.Delete("profile/:id", controllers.DeleteUserProfile)
	app.Get("profile/show", controllers.ShowProfiles)


	app.Post("/cart/add", controllers.AddToCart)
	app.Get("/cart/:id", controllers.GetCartItems)

}	
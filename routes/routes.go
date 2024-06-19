package routes

import (
	"github.com/InfamousFreak/Tech-Task-24/controllers"
	"github.com/gofiber/fiber/v2"
	//"github.com/InfamousFreak/Tech-Task-24/middlewares"
)

func SetupRouter(app *fiber.App) {

	//app.Use(AuthMiddleware(config.Secret))

	app.Post("/profile/create", controllers.CreateUserProfile)



	app.Post("/profile/:id", controllers.UpdateUserProfile)
	app.Delete("/profile/:id", controllers.DeleteUserProfile)

	app.Post("/profile/selectrole", controllers.SelectRole)

	app.Get("/menu/get", controllers.GetMenuItems)
	app.Post("/menu/create", controllers.CreateMenuItem)
	app.Get("/menu/search", controllers.SearchMenuItemsByTags) //new search route for tags

	app.Post("/cart/add", controllers.AddToCart)
	app.Get("/cart/:id", controllers.GetCartItems)

	app.Get("/profile/show", controllers.ShowProfiles)

}	
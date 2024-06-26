package routes

import (
    "github.com/InfamousFreak/Tech-Task-24/controllers"
    "github.com/gofiber/fiber/v2"
    "github.com/InfamousFreak/Tech-Task-24/middlewares"
)

func SetupRouter(app *fiber.App) {
    jwt := middlewares.AuthMiddleware()  

    app.Post("/profile/create", controllers.CreateUserProfile)
    app.Post("/profile/:id", jwt, controllers.UpdateUserProfile)
    app.Delete("/profile/:id", jwt, controllers.DeleteUserProfile)
    app.Get("/profile/:id/details", jwt, controllers.ShowUserProfile)
    app.Get("/profile/show", jwt, controllers.ShowProfiles)

    app.Get("/menu/get", controllers.GetMenuItems)
    app.Post("/menu/create", jwt, controllers.CreateMenuItem)
    app.Get("/menu/search", jwt, controllers.SearchMenuItemsByTags)
    app.Patch("/menu/:id", controllers.UpdateMenuItem)

    app.Post("/cart/add", jwt, controllers.AddToCart)
    app.Get("/cart/:user_id", controllers.GetCartItems)


}
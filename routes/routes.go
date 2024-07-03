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
    app.Delete("/profile/:id", jwt, controllers.DeleteUserProfile) //user + admin privilege
    app.Get("/profile/:id/details", jwt, controllers.ShowUserProfile) //admin privilege


    app.Get("/menu/get", controllers.GetMenuItems)
    app.Post("/menu/create", jwt, controllers.CreateMenuItem) //admin privilege
    app.Get("/menu/search", jwt, controllers.SearchMenuItemsByTags)
    app.Patch("/menu/:id", controllers.UpdateMenuItem) //admin privilege

    app.Post("/cart/add", jwt, controllers.AddToCart)
    app.Get("/cart/:user_id", controllers.GetCartItems)
    app.Put("/cart/upsert", controllers.UpsertCartItem) 
    app.Delete("/cart/delete", controllers.DeleteCartItem)


    app.Post("/admin/signup", controllers.AdminSignup)
    app.Get("/admin/profiles", controllers.GetAllAdminProfiles)
    app.Get("/profile/show", jwt, controllers.ShowProfiles) //admin privileges
    app.Delete("/admin/delete/:id", controllers.DeleteAdminProfile)
    app.Delete("/admin/cart", controllers.AdminDeleteUserCartItem)
}
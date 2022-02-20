package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/controllers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", config.AuthMiddleware) // /api

	groups := api.Group("/groups")
	groups.Get("/", controllers.GetGroups)

	users := api.Group("/users")
	users.Get("/current", controllers.GetCurrentUser)
	users.Post("/signup", controllers.SignUpUser)
	users.Post("/profile/upload", controllers.UploadProfile)

	friends := users.Group("/friends")
	friends.Post("/req", controllers.RequestFriend)
	friends.Post("/act", controllers.AcceptFriendRequest)
}

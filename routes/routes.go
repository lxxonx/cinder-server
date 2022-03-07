package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/controllers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api") // /api

	users := api.Group("/users")
	users.Get("/current", config.AuthMiddleware, controllers.GetCurrentUser)
	users.Get("/", controllers.GetAllUsers)
	users.Post("/signup", controllers.SignUpUser)
	users.Post("/pic", config.AuthMiddleware, controllers.UploadProfile)
	users.Post("/sc", config.AuthMiddleware, controllers.UploadStudentCard)

	friends := api.Group("/friends")
	friends.Get("/", config.AuthMiddleware, controllers.GetFriends)
	friends.Get("/req", config.AuthMiddleware, controllers.GetFriendRequest)
	friends.Post("/req", config.AuthMiddleware, controllers.RequestFriend)
	friends.Post("/act", config.AuthMiddleware, controllers.AcceptFriendRequest)
	friends.Post("/search", controllers.SearchFriend)

	groups := api.Group("/groups")
	groups.Get("/", config.AuthMiddleware, controllers.GetGroups)
	groups.Get("/my", config.AuthMiddleware, controllers.GetMyGroups)
	groups.Post("/join", config.AuthMiddleware, controllers.JoinGroup)
	groups.Post("/pic", config.AuthMiddleware, controllers.UploadGroupProfile)
}

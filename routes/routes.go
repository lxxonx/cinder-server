package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/controllers"
)


func Setup(app *fiber.App) {
	
	api := app.Group("/api")      // /api

	api.Post("/photo", controllers.UploadPhoto)
}

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/routes"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Use(config.SetupFirebase)
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

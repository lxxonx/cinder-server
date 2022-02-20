package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/routes"
)

func main() {

	err := godotenv.Load(".env")
	port := os.Getenv("PORT")
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(config.SetupFirebase)
	routes.SetupRoutes(app)

	app.Listen(":" + port)
}

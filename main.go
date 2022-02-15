package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lxxonx/cinder-server/routes"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    app := fiber.New()

    routes.Setup(app)

    app.Listen(":3000")
}
package config

import (
	"context"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	env := os.Getenv("ENV")

	if env == "dev" {
		c.Locals("uid", "Zjx9cgDlESMb9Nq4WcuDwNe4gSu1")
		return c.Next()
	}

	headers := c.GetReqHeaders()
	jwt := strings.Split(headers["Authorization"], " ")[1]
	firebaseApp := c.Locals("firebase").(*firebase.App)

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to authenticate",
		})
	}

	token, err := auth.VerifyIDToken(context.Background(), jwt)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"ok":      false,
			"message": "Invalid token",
		})
	}

	if token == nil {
		return c.Status(401).JSON(fiber.Map{
			"ok":      false,
			"message": "Invalid token",
		})
	}
	c.Locals("uid", token.UID)

	return c.Next()
}

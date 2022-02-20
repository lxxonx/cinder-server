package config

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/dto"
)

func AuthMiddleware(c *fiber.Ctx) error {
	firebaseApp := c.Locals("firebase").(*firebase.App)
	input := new(dto.AuthInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to authenticate",
		})
	}

	token, err := auth.VerifyIDToken(context.Background(), input.Token)
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
	c.Locals("user", token.UID)

	return c.Next()
}

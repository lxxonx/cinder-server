package config

import (
	"context"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/models"
)

func AuthMiddleware(c *fiber.Ctx) error {
	env := os.Getenv("ENV")

	if env == "dev" {
		c.Locals("user", models.UserCtx{Uid: "Zjx9cgDlESMb9Nq4WcuDwNe4gSu1", Username: "wleifns"})
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
	c.Locals("user", models.UserCtx{Uid: token.UID, Username: token.Issuer})

	return c.Next()
}

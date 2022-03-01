package config

import (
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/ent/user"
)

func AuthMiddleware(c *fiber.Ctx) error {
	env := os.Getenv("ENV")

	if env == "dev" {
		c.Locals("userId", 1)
		return c.Next()
	}

	headers := c.GetReqHeaders()
	jwt := strings.Split(headers["Authorization"], " ")[1]
	firebaseApp := c.Locals("firebase").(*firebase.App)

	auth, err := firebaseApp.Auth(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to authenticate",
		})
	}

	token, err := auth.VerifyIDToken(c.Context(), jwt)
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

	userId := DB.User.Query().Where(user.UID(token.UID)).OnlyIDX(c.Context())

	c.Locals("userId", userId)
	return c.Next()
}

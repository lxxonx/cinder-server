package config

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()

	if headers["Authorization"] == "" {
		return c.Status(401).JSON(fiber.Map{
			"ok":      false,
			"message": "Unauthorized",
		})
	}

	jwt := strings.Split(headers["Authorization"], " ")[1]
	if jwt == "" {
		return c.Status(401).JSON(fiber.Map{
			"ok":      false,
			"message": "Unauthorized",
		})
	}

	auth, err := FBA.Auth(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to authenticate",
		})
	}

	token, err := auth.VerifyIDToken(c.Context(), jwt)
	if err != nil {
		fmt.Println(err.Error())
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

	c.Locals("userId", token.UID)
	return c.Next()
}

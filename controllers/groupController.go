package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetGroups(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
	})
}

func CreateGroup(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
	})
}

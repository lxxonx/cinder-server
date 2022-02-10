package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadPhoto(c *fiber.Ctx) error {
	// Get first file from form field "document":
	file, err := c.FormFile("document")

	if err != nil {
		return err
	} else {
		// 👷 Save file to root directory:
		c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		// 👷 Save file inside uploads folder under current working directory:
		c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		// 👷 Save file using a relative path:
		c.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))
		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
		})
	}
}
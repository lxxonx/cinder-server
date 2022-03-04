package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/ent/user"
)

func GetChats(c *fiber.Ctx) error {

	userId := c.Locals("userId").(string)

	input := new(dto.GetChatsInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	chats := config.DB.User.Query().Where(user.IDEQ(userId)).QueryChatroom().AllX(c.Context())

	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "success",
		"data": fiber.Map{
			"chats": chats,
		},
	})
}

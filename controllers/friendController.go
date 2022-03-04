package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/ent/user"
	"github.com/lxxonx/cinder-server/models"
)

func GetFriendRequest(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var reqs []models.User
	config.DB.User.Query().
		Where(user.IDEQ(userId)).
		QueryFriendsReq().
		Select(user.FieldID, user.FieldUsername, user.FieldBio, user.FieldUni, user.FieldDep).
		ScanX(c.Context(), &reqs)

	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "success",
		"data": fiber.Map{
			"requests": reqs,
		},
	})
}

func GetFriends(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var reqs []models.User
	config.DB.User.Query().
		Where(user.IDEQ(userId)).
		QueryFriends().
		Select(user.FieldID, user.FieldUsername, user.FieldBio, user.FieldUni, user.FieldDep).
		ScanX(c.Context(), &reqs)

	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "success",
		"data": fiber.Map{
			"friends": reqs,
		},
	})
}

func RequestFriend(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	input := new(dto.RequestFriendInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	me := config.DB.User.Query().Where(user.IDEQ(userId)).OnlyX(c.Context())

	friend, err := config.DB.User.Query().Where(user.IDEQ(input.FID)).Only(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to find friend",
		})
	}

	me.Update().AddRequests(friend).ExecX(c.Context())

	return c.Status(201).JSON(fiber.Map{
		"ok":      true,
		"message": "Successfully sent friend request",
	})
}

func AcceptFriendRequest(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	input := new(dto.RequestFriendInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	me := config.DB.User.Query().Where(user.IDEQ(userId)).OnlyX(c.Context())
	_, err := me.QueryFriendsReq().Where(user.IDEQ(input.FID)).Exist(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Invalid request",
		})
	}
	friend := config.DB.User.Query().Where(user.IDEQ(input.FID)).OnlyX(c.Context())

	err = me.Update().AddFriends(friend).RemoveFriendsReq(friend).Exec(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "success",
	})
}

func DeleteFriendRequest(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	input := new(dto.RequestFriendInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	me := config.DB.User.Query().Where(user.IDEQ(userId)).OnlyX(c.Context())
	_, err := me.QueryFriendsReq().Where(user.IDEQ(input.FID)).Exist(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Invalid request",
		})
	}
	friend := config.DB.User.Query().Where(user.IDEQ(input.FID)).OnlyX(c.Context())

	me.Update().RemoveFriendsReq(friend).ExecX(c.Context())

	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "success",
	})
}

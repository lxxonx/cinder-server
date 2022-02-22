package controllers

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/models"
)

func GetGroups(c *fiber.Ctx) error {
	input := new(dto.GetGroupsInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
	})
}

func CreateGroup(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserCtx)

	input := new(dto.CreateGroupIput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	firebaseApp := c.Locals("firebase").(*firebase.App)
	db, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to connect to database",
		})
	}
	uid := uuid.New()
	group := db.Collection("groups").Doc(uid.String())

	for _, member := range input.Friends {
		fmt.Print(member)
		friend, err := db.Collection("users").Doc(user.Uid).Collection("friends").Where("username", "==", member).Documents(context.Background()).GetAll()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"ok":      false,
				"message": err.Error(),
			})
		}
		if len(friend) == 0 || friend[0].Data() == nil {
			return c.Status(500).JSON(fiber.Map{
				"ok":      false,
				"message": "Friend not found",
			})
		}
		fmt.Print()

		f, err := db.Collection("users").Doc(friend[0].Data()["uid"].(string)).Get(context.Background())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"ok":      false,
				"message": err.Error(),
			})
		}
		group.Collection("members").Doc(friend[0].Data()["uid"].(string)).Set(context.Background(), f)
	}
	me, err := db.Collection("users").Doc(user.Uid).Get(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}
	group.Collection("members").Doc(user.Uid).Set(context.Background(), me)
	_, err = group.Set(context.Background(), map[string]interface{}{
		"uid":       uid.String(),
		"groupName": input.GroupName,
		"pics":      input.Pics,
		"bio":       input.Bio,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
	})
}

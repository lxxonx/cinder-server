package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/models"
)

func RequestFriend(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserCtx)
	input := new(dto.RequestFriendInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	db, err := c.Locals("firebase").(*firebase.App).Firestore(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Fail to connect to database",
		})
	}

	frRef, err := db.Collection("users").Where("username", "==", input.FriendName).Documents(context.Background()).GetAll()

	if err != nil || len(frRef) == 0 {
		fmt.Print(err)
		return c.Status(200).JSON(fiber.Map{
			"ok":      false,
			"message": "Friend not found",
		})
	} else if len(frRef) > 1 {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Multiple friends found",
		})
	}

	// friendsUid := frRef[0].Data()["uid"].(string)
	friendsReq, err := frRef[0].Ref.Collection("friendsReq").Doc(user.Uid).Get(context.Background())
	if friendsReq.Data() != nil {
		return c.Status(200).JSON(fiber.Map{
			"ok":      false,
			"message": "Friend request already sent",
		})
	}
	if err != nil {
		if strings.Split(err.Error(), " ")[4] == "NotFound" {
			_, err = frRef[0].Ref.Collection("friendsReq").Doc(user.Uid).Set(context.Background(), map[string]interface{}{
				"uid":      user.Uid,
				"username": user.Username,
			}, firestore.MergeAll)
			if err != nil {
				fmt.Print(err)
				return c.Status(500).JSON(fiber.Map{
					"ok":      false,
					"message": "Unable to connect to database",
				})
			}
		}

	}

	return c.Status(201).JSON(fiber.Map{
		"ok":      true,
		"message": "Successfully sent friend request",
	})
}

func AcceptFriendRequest(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserCtx)
	input := new(dto.RequestFriendInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	db, err := c.Locals("firebase").(*firebase.App).Firestore(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Fail to connect to database",
		})
	}

	reqRef, err := db.Collection("users").Doc(user.Uid).Collection("friendsReq").Where("username", "==", input.FriendName).Documents(context.Background()).GetAll()
	if err != nil {
		fmt.Print(err)
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}
	if len(reqRef) == 0 || reqRef[0].Data() == nil {
		return c.Status(200).JSON(fiber.Map{
			"ok":      false,
			"message": "cannot find friend request",
		})
	}
	_, err = db.Collection("users").Doc(user.Uid).Collection("friendsReq").Doc(reqRef[0].Data()["uid"].(string)).Delete(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "fail to delete friend request",
		})
	}
	friend, _ := db.Collection("users").Doc(user.Uid).Collection("friends").Doc(reqRef[0].Data()["uid"].(string)).Get(context.Background())
	if friend.Data() != nil {
		return c.Status(200).JSON(fiber.Map{
			"ok":      false,
			"message": "friend already added",
		})
	}
	_, err = db.Collection("users").Doc(user.Uid).Collection("friends").Doc(reqRef[0].Data()["uid"].(string)).Set(context.Background(), map[string]interface{}{
		"Ref":        "users/" + reqRef[0].Data()["uid"].(string),
		"username":   reqRef[0].Data()["username"],
		"CreateTime": time.Now(),
		"UpdateTime": time.Now(),
		"ReadTime":   time.Now(),
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "friend request not found",
		})
	}

	_, err = db.Collection("users").Doc(reqRef[0].Data()["uid"].(string)).Collection("friends").Doc(user.Uid).Set(context.Background(), map[string]interface{}{
		"Ref":        "users/" + user.Uid,
		"username":   user.Username,
		"CreateTime": time.Now(),
		"UpdateTime": time.Now(),
		"ReadTime":   time.Now(),
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "friend request not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "Successfully accepted friend request"})
}

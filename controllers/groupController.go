package controllers

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/models"
)

func GetGroups(c *fiber.Ctx) error {

	input := new(dto.GetGroupsInput)

	if len(c.Body()) > 0 {
		if err := c.BodyParser(input); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"ok":      false,
				"message": "Unable to parse request body",
			})
		}
	} else {
		input.Limit = 10
		input.Cursor = ""
	}

	firebaseApp := c.Locals("firebase").(*firebase.App)
	db, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to connect to database",
		})
	}
	groups, err := db.Collection("groups").OrderBy("uid", firestore.Desc).Limit(input.Limit).Documents(context.Background()).GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(200).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to connect to database",
		})
	}

	var res []models.Group

	for _, group := range groups {
		users, err := group.Ref.Collection("members").Documents(context.Background()).GetAll()
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(200).JSON(fiber.Map{
				"ok":      false,
				"message": "Unable to connect to database",
			})
		}
		var members []models.User
		for _, u := range users {
			user, err := db.Collection("users").Doc(u.Ref.ID).Get(context.Background())
			if err != nil {
				fmt.Println(err.Error())
				return c.Status(200).JSON(fiber.Map{
					"ok":      false,
					"message": "Unable to connect to database",
				})
			}
			member := models.User{}
			member = member.Create(user)
			members = append(members, member)
		}
		model := models.Group{
			Uid:        group.Data()["uid"].(string),
			GroupName:  group.Data()["groupName"].(string),
			Pics:       group.Data()["pics"].([]interface{}),
			CreateTime: group.CreateTime,
			Bio:        group.Data()["bio"].(string),
			Members:    members,
		}
		res = append(res, model)
	}
	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Fetch Groups successfully",
		"data": fiber.Map{
			"groups": res,
		},
	})
}

func CreateGroup(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserCtx)

	input := new(dto.CreateGroupInput)
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

	// TODO
	// Check if user is in group
	// Duplication check
	for _, member := range input.Friends {
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
		"uid":        uid.String(),
		"groupName":  input.GroupName,
		"pics":       input.Pics,
		"bio":        input.Bio,
		"CreateTime": time.Now(),
		"UpdateTime": time.Now(),
		"ReadTime":   time.Now(),
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

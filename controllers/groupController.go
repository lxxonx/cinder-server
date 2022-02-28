package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/ent/group"
	"github.com/lxxonx/cinder-server/ent/user"
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

	groups := config.DB.Group.
		Query().
		Where(
			group.ReadAtGTE(time.Now().AddDate(0, 0, -2)),
		).
		QueryMembers().
		QueryPics()

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Fetch Groups successfully",
		"data": fiber.Map{
			"groups": groups,
		},
	})
}

func CreateGroup(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)

	input := new(dto.CreateGroupInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	var friendsIDs []string

	friendsIDs = append(friendsIDs, uid)
	for _, friend := range input.Friends {
		fr, _ := config.DB.User.Query().Where(user.Username(friend)).First(c.Context())

		friendsIDs = append(friendsIDs, fr.ID)
	}

	gid := uuid.New().String()

	config.DB.Group.
		Create().
		SetID(gid).SetGroupname(input.GroupName).
		SetBio(input.Bio).AddMemberIDs(friendsIDs...).SaveX(c.Context())

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
	})
}

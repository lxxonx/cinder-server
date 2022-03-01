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
	userId := c.Locals("userId").(int)

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

	gid := config.DB.User.Query().Where(user.IDEQ(userId)).QueryGroup().OnlyIDX(c.Context())

	groups := config.DB.Group.Query().Where(
		group.And(
			group.ReadAtGTE(time.Now().AddDate(0, 0, -2)),
			group.IDNEQ(gid),
		),
	).WithPics().WithMembers().AllX(c.Context())

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Fetch Groups successfully",
		"data": fiber.Map{
			"groups": groups,
		},
	})
}

func JoinGroup(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)

	input := new(dto.CreateGroupInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	var friendsIDs []int

	friendsIDs = append(friendsIDs, userId)
	for _, friend := range input.Friends {
		frID, err := config.DB.User.Query().Where(user.IDEQ(userId)).Clone().QueryFriends().Where(user.UsernameEQ(friend)).OnlyID(c.Context())
		if err != nil {
			return c.Status(200).JSON(fiber.Map{
				"ok":      false,
				"message": "cannot find friend",
			})
		}
		friendsIDs = append(friendsIDs, frID)
	}

	gid := uuid.New().String()

	g := config.DB.Group.
		Create().
		SetUID(gid).SetGroupname(input.GroupName).
		SetBio(input.Bio).AddMemberIDs(friendsIDs...).SaveX(c.Context())

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
		"data": fiber.Map{
			"group": g,
		},
	})
}

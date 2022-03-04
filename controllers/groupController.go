package controllers

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/anthonynsimon/bild/blur"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/ent/group"
	"github.com/lxxonx/cinder-server/ent/user"
)

func GetGroups(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

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

	gid, err := config.DB.User.Query().Where(user.IDEQ(userId)).QueryGroup().OnlyID(c.Context())

	if err != nil {
		groups := config.DB.Group.Query().Where(
			group.And(
				group.ReadAtGTE(time.Now().AddDate(0, 0, -2)),
			),
		).WithPics().WithMembers().AllX(c.Context())
		return c.JSON(fiber.Map{
			"ok":      true,
			"message": "Fetch Groups successfully",
			"data": fiber.Map{
				"groups": groups,
			},
		})
	} else {
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

}

func JoinGroup(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	input := new(dto.CreateGroupInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	var friendsIDs []string

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

	gid := uuid.New()

	g := config.DB.Group.
		Create().
		SetID(gid).SetGroupname(input.GroupName).
		SetBio(input.Bio).AddMemberIDs(friendsIDs...).SaveX(c.Context())

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create Group successfully",
		"data": fiber.Map{
			"group": g,
		},
	})
}

func UploadGroupProfile(c *fiber.Ctx) error {
	firebaseApp := c.Locals("firebase").(*firebase.App)

	handler, err := c.FormFile("photo")
	if err != nil {
		return err
	}
	file, err := handler.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	var original image.Image
	switch strings.Split(handler.Filename, ".")[1] {
	case "png":
		original, err = png.Decode(file)
	case "jpeg", "jpg":
		original, err = jpeg.Decode(file)
	}
	if err != nil {
		return c.JSON(fiber.Map{
			"ok":      false,
			"message": "File is not an image",
		})
	}
	var pics []image.Image
	pics = append(pics, original)

	client, err := firebaseApp.Storage(context.Background())
	if err != nil {
		panic(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		panic(err)
	}

	for i := 0.0; i < 3; i++ {
		nb := blur.Gaussian(original, (i+1.0)*20.0)

		pics = append(pics, nb)
	}

	id := uuid.New()
	for i, pic := range pics {

		fileName := id.String() + "/p" + fmt.Sprintf("%d", i)

		writer := bucket.Object(fileName).NewWriter(context.Background())
		// writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
		defer writer.Close()
		buf := new(bytes.Buffer)
		if err = png.Encode(buf, pic); err != nil {
			return c.JSON(fiber.Map{
				"ok":      false,
				"message": "Can't get file",
			})
		}

		_, err = io.Copy(writer, bytes.NewReader(buf.Bytes()))
		if err != nil {
			return c.JSON(fiber.Map{
				"ok":      false,
				"message": "Can't get file",
			})
		}
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "File uploaded successfully",
	})

}

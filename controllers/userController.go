package controllers

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"io"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/anthonynsimon/bild/blur"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/dto"
)

func SignUpUser(c *fiber.Ctx) error {
	input := new(dto.SignUpInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}

	firebaseApp := c.Locals("firebase").(*firebase.App)
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to connect to auth",
		})
	}

	params := (&auth.UserToCreate{}).
		Email(input.Email).
		EmailVerified(false).
		Password(input.Password).
		DisplayName(input.Username).
		Disabled(false)
	user, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Print(err)

		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Fail to create user",
		})
	}

	db, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to connect to database",
		})
	}

	_, err = db.Collection("users").Doc(user.UID).Set(context.Background(), map[string]interface{}{
		"uid":      user.UID,
		"username": input.Username,
		"email":    input.Email,
		"password": input.Password,
		"avatar":   "",
		"dep":      input.Dep,
		"uni":      input.Uni,
		"bio":      "",
	})
	if err != nil {
		fmt.Print(err)
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Fail to create user",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Create user successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	firebaseApp := c.Locals("firebase").(*firebase.App)

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to authenticate",
		})
	}

	token, err := auth.CustomToken(context.Background(), "some-uid")
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}
	fmt.Print(token)

	return c.Next()
}

func GetCurrentUser(c *fiber.Ctx) error {
	uid := c.Locals("user")
	if uid == nil {
		return c.Status(401).JSON(fiber.Map{
			"ok":      false,
			"message": "Not logged in",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": uid,
	})
}

func UploadProfile(c *fiber.Ctx) error {
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

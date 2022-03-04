package controllers

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"io"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/anthonynsimon/bild/blur"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/config"
	"github.com/lxxonx/cinder-server/dto"
	"github.com/lxxonx/cinder-server/ent/user"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	input := new(dto.SignUpInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "Unable to parse request body",
		})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "unable to hash password",
		})
	}

	_, err = config.DB.User. // UserClient.
					Create().         // User create builder.
					SetID(input.Uid). // Set uid.
					SetUsername(input.Username).
					SetActualName(input.ActualName).
					SetBirthYear(input.BirthYear).
					SetGender(input.Gender).
					SetPassword(password).
					SetBio("").
					SetUni(input.Uni).
					SetDep(input.Dep).
					Save(c.Context()) // Create and return.
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
	userId := c.Locals("userId")
	if userId == nil {
		return c.Status(401).JSON(fiber.Map{
			"ok":      false,
			"message": "Not logged in",
		})
	}

	me := config.DB.User.Query().Where(user.IDEQ(userId.(string))).OnlyX(c.Context())

	return c.Status(200).JSON(fiber.Map{
		"ok": true,
		"data": fiber.Map{
			"me": me,
		},
	})
}
func GetAllUsers(c *fiber.Ctx) error {

	u, err := config.DB.User.Query().
		Select(
			user.FieldUsername,
			user.FieldID,
			user.FieldBio,
			user.FieldUni,
			user.FieldDep,
		).
		WithPics().All(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"ok": true,
		"data": fiber.Map{
			"users": u,
		},
	})
}

func UploadProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	firebaseApp := config.FBA

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
	storage_address := os.Getenv("FIREBASE_BUCKET_ADDRESS")
	url := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", storage_address, userId+"/p0")

	_, err = config.DB.Pic.Create().SetUserID(userId).SetURL(url).SetID(id).Save(c.Context())
	if err != nil {
		return c.JSON(fiber.Map{
			"ok":      false,
			"message": "can't save pic",
		})
	}
	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "File uploaded successfully",
		"data": fiber.Map{
			"url": url,
		},
	})

}

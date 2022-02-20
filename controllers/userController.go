package controllers

import (
	"bytes"
	"context"
	"fmt"

	"io"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
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
	return c.Next()
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
			"message": "Invalid token",
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
	// original, err := jpeg.Decode(file)

	if err != nil {
		return err
	} else {

		client, err := firebaseApp.Storage(context.Background())
		if err != nil {
			panic(err)
		}

		bucket, err := client.DefaultBucket()
		if err != nil {
			panic(err)
		}
		id := uuid.New()
		fileName := id.String() + "/profile.jpg"

		object := bucket.Object(fileName)
		writer := object.NewWriter(context.Background())
		// writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
		defer writer.Close()
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
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

		storage_address := os.Getenv("FIREBASE_BUCKET_ADDRESS")

		file_url := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", storage_address, fileName)

		// blur3 := blur.Gaussian(original, 40.0)
		// image3, err := os.Create("blur.png")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer image3.Close()

		// err = png.Encode(image3, blur3)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// writer = bucket.Object("/p/blur3.png").NewWriter(context.Background())
		// _, err = io.Copy(writer, image3)

		// if err != nil {
		// 	return c.JSON(fiber.Map{
		// 		"ok":      false,
		// 		"message": "Can't upload file",
		// 	})
		// }

		return c.JSON(fiber.Map{
			"ok":      true,
			"message": "File uploaded successfully",
			"data": fiber.Map{
				"file_url": file_url,
			},
		})
	}
}

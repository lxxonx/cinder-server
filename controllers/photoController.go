package controllers

import (
	"fmt"
	"io"
	"log"
	"os"

	fb "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

func UploadPhoto(c *fiber.Ctx) error {
	// Get first file from form field "document":
	handler, err := c.FormFile("document")
	if err != nil{
		return err
	}
	file, err := handler.Open()
	if err != nil{
		return err
	}
	defer file.Close()
	if err != nil {
		return err
	} else {
		// 👷 Save file to root directory:
		// c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		// 👷 Save file inside uploads folder under current working directory:
		// c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		// 👷 Save file using a relative path:
		// c.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))

		// image, err := imgio.Open(file.Filename);
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// box := blur.Box(image, 10.0)
		// if err := imgio.Save("box.png", box, imgio.PNGEncoder()); err != nil {
		// fmt.Println(err)

		storage_address := os.Getenv("FIREBASE_BUCKET_ADDRESS")
		project_id := os.Getenv("FIREBASE_PROJECT_ID")
		config := &fb.Config{
			StorageBucket: storage_address,
			ProjectID: project_id,
		}
		id := uuid.New() 

		opt := option.WithCredentialsFile("ServiceAccountKey.json")
		app, err := fb.NewApp(c.Context(), config, opt)
		if err != nil {
			log.Fatalln(err)
		}
		
		client, err := app.Storage(c.Context())
		if err != nil {
			log.Fatalln(err)
		}
	
		bucket, err := client.DefaultBucket()
		if err != nil {
			log.Fatalln(err)
		}
		fileName := "p/profile_" + id.String() + ".png"

		object := bucket.Object(fileName)
		writer := object.NewWriter(c.Context())
        writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()} 
		defer writer.Close()

		byteSize, err := io.Copy(writer, file)
		if err != nil {
			return c.JSON(fiber.Map{
				"ok": false,
				"message": "File uploaded successfully",
			})
		}
		fmt.Printf("File size uploaded: %v\n", byteSize)

		// if err := object.ACL().Set(c.Context(), storage.AllUsers, storage.RoleReader); err != nil {
		// 	return err
		// }
		uploadedImageUrl := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", storage_address, fileName)
		fmt.Printf("Image uploaded successfully: %v", uploadedImageUrl)


		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
		})
	}
}
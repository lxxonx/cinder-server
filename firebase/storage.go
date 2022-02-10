package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
	storage_address := os.Getenv("FIREBASE_BUCKET_ADDRESS")
	config := &firebase.Config{
        StorageBucket: storage_address,
	}
	
	opt := option.WithCredentialsFile("ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
        log.Fatalln(err)
	}
	
	return app
}

func AddFile(app *firebase.App) {
	client, err := app.Storage(context.Background())
	if err != nil {
        log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
        log.Fatalln(err)
	}

	println(bucket)
}


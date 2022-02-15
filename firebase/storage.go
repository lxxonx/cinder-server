package firebase

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitStorage() *storage.BucketHandle {
	storage_address := os.Getenv("FIREBASE_BUCKET_ADDRESS")
	config := &firebase.Config{
        StorageBucket: storage_address,
	}
	
	opt := option.WithCredentialsFile("ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
        log.Fatalln(err)
	}
	
	client, err := app.Storage(context.Background())
	if err != nil {
        log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()

	if err != nil {
        log.Fatalln(err)
	}

	return bucket
}



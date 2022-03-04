package config

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FBA *firebase.App

func SetupFirebase() {
	storage_address := os.Getenv("FIREBASE_BUCKET_ADDRESS")
	project_id := os.Getenv("FIREBASE_PROJECT_ID")
	config := &firebase.Config{
		StorageBucket: storage_address,
		ProjectID:     project_id,
	}

	opt := option.WithCredentialsFile("ServiceAccountKey.json")

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic("Firebase load error")
	}

	FBA = app
}

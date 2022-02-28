package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lxxonx/cinder-server/ent"
)

var DB *ent.Client

func DBConnect() {
	db := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))

	client, err := ent.Open("postgres", db)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	DB = client
}

package config

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

var GCSClient *storage.Client

func InitGCS(ctx context.Context) {
	if !AppConfig.GCSEnabled {
		log.Println("GCS is disabled. Skipping client initialization.")
		return
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize GCS client: %v", err)
	}

	GCSClient = client
	log.Println("Connected to Google Cloud Storage.")
}

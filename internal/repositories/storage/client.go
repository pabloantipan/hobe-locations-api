package storage

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/pabloantipan/hobe-locations-api/config"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"google.golang.org/api/option"
)

func NewStorageClient(cfg *config.Config) *storage.Client {
	if cfg.StorageServiceAccountPath == "" {
		log.Printf("DATASTORE_SERVICE_ACCOUNT_PATH environment variable not set")
		return nil
	}

	data, err := os.ReadFile(cfg.StorageServiceAccountPath)
	if err != nil {
		log.Printf("Failed to read service account file: %v", err)
		return nil
	}

	var sa models.ServiceAccount
	if err := json.Unmarshal(data, &sa); err != nil {
		log.Printf("Failed to unmarshal service account file: %v", err)
		return nil
	}

	ctx := context.Background()
	client, err := storage.NewClient(
		ctx,
		option.WithCredentialsFile(cfg.StorageServiceAccountPath),
	)
	if err != nil {
		log.Fatalf("failed to create storage client: %v", err)
		return nil
	}

	return client
}

func CloseStorageClient(client *storage.Client) {
	if err := client.Close(); err != nil {
		log.Printf("Failed to close storage client: %v", err)
	}
}

package datastore

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/pabloantipan/hobe-locations-api/config"
	"github.com/pabloantipan/hobe-locations-api/internal/models"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

func NewDatastoreClient(cfg *config.Config) *datastore.Client {
	if cfg.DatastoreServiceAccountPath == "" {
		log.Printf("DATASTORE_SERVICE_ACCOUNT_PATH environment variable not set")
		return nil
	}

	data, err := os.ReadFile(cfg.DatastoreServiceAccountPath)
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
	client, err := datastore.NewClientWithDatabase(
		ctx,
		sa.ProjectID,
		cfg.DatastoreDatabaseID,
		option.WithCredentialsFile(cfg.DatastoreServiceAccountPath),
	)

	if err != nil {
		log.Fatalf("Failed to create datastore client: %v", err)
		return nil
	}

	return client
}

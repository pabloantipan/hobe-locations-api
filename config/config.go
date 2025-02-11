package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Who                         string
	Version                     string
	Port                        string
	ProjectID                   string
	DatastoreDatabaseID         string
	DatastoreServiceAccountPath string
	LoggingServiceAccountPath   string
	StorageServiceAccountPath   string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		Who:                         "[" + os.Getenv("WHO") + "]",
		Version:                     os.Getenv("VERSION"),
		Port:                        os.Getenv("PORT"),
		ProjectID:                   os.Getenv("PROJECT_ID"),
		DatastoreDatabaseID:         os.Getenv("DATASTORE_DATABASE_ID"),
		DatastoreServiceAccountPath: os.Getenv("DATASTORE_SERVICE_ACCOUNT_FILE"),
		LoggingServiceAccountPath:   os.Getenv("LOGGING_SERVICE_ACCOUNT_FILE"),
		StorageServiceAccountPath:   os.Getenv("STORAGE_SERVICE_ACCOUNT_FILE"),
	}

	return config, nil
}

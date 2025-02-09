package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version                     string
	Port                        string
	ProjectID                   string
	DatastoreServiceAccountPath string
	CloudLoggingCredentialsFile string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		Version:                     os.Getenv("VERSION"),
		Port:                        os.Getenv("PORT"),
		ProjectID:                   os.Getenv("PROJECT_ID"),
		DatastoreServiceAccountPath: os.Getenv("DATASTORE_SERVICE_ACCOUNT_FILE"),
		CloudLoggingCredentialsFile: os.Getenv("LOGGING_SERVICE_ACCOUNT_FILE"),
	}

	return config, nil
}

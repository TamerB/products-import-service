package config

import (
	"os"

	"github.com/TamerB/products-import-service/constants"
)

type DBConfig struct {
	DBDriver string
	DBSource string
	URL      string
}

func NewConfig() *DBConfig {
	baseUrl, ok := os.LookupEnv(constants.EnvURL)
	if !ok {
		baseUrl = "localhost:8080"
	}
	return &DBConfig{
		DBDriver: os.Getenv(constants.EnvDBDriver),
		DBSource: os.Getenv(constants.EnvDBSource),
		URL:      baseUrl,
	}
}

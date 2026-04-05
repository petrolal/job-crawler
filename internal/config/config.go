// Package config loads application configuration from environment variables.
package config

import (
	"os"
	"strconv"
)

// Config holds all application configuration values.
type Config struct {
	AdzunaAppID   string
	AdzunaAPIKey  string
	AdzunaPages   int
	AdzunaPerPage int
}

// Load reads configuration from environment variables and returns a Config.
// Panics if required numeric variables are missing or malformed.
func Load() Config {
	pages, err := strconv.Atoi(os.Getenv("ADZUNA_PAGES"))
	if err != nil {
		panic(err)
	}

	perPage, err := strconv.Atoi(os.Getenv("ADZUNA_PER_PAGE"))
	if err != nil {
		panic(err)
	}

	return Config{
		AdzunaAppID:   os.Getenv("ADZUNA_APP_ID"),
		AdzunaAPIKey:  os.Getenv("ADZUNA_API_KEY"),
		AdzunaPages:   pages,
		AdzunaPerPage: perPage,
	}
}

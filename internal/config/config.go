package config

import "os"

type Config struct {
	AdzunaAppID  string
	AdzunaAPIKey string
}

func Load() Config {
	return Config{
		AdzunaAppID:  os.Getenv("ADZUNA_APP_ID"),
		AdzunaAPIKey: os.Getenv("ADZUNA_API_KEY"),
	}
}

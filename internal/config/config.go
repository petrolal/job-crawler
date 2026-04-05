package config

import (
	"os"
	"strconv"
)

type Config struct {
	AdzunaAppID   string
	AdzunaAPIKey  string
	AdzunaPages   int
	AdzunaPerPage int
}

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

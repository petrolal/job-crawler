package main

import (
	"fmt"
	"log"
	"os"

	"jobs-crawler/internal/api"
	"jobs-crawler/internal/config"
	"jobs-crawler/internal/service"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
	"jobs-crawler/internal/sources/lever"
	"jobs-crawler/internal/store"
)

func main() {
	c := config.Load()

	crawler := service.CrawlerService{
		Adzuna: adzuna.Client{
			AppID:       c.AdzunaAppID,
			APIKey:      c.AdzunaAPIKey,
			Pages:       c.AdzunaPages,
			ResultsPage: c.AdzunaPerPage,
		},
		Greenhouse: []greenhouse.Client{
			{Company: "nubank"},
			{Company: "vtex"},
			{Company: "ifood"},
			{Company: "gympass"},
			{Company: "quintoandar"},
		},
		Lever: []lever.Client{
			{Company: "cloudwalk"},
		},
	}

	jobStore := store.NewJobStore()

	go func() {
		jobs, err := crawler.Run()
		if err != nil {
			log.Printf("crawler error: %v", err)
		}
		jobStore.SetJobs(jobs)
		fmt.Printf("\n✅ Total QA jobs: %d\n", len(jobs))
	}()

	r := api.NewRouter(jobStore)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

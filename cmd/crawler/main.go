package main

import (
	"fmt"
	"log"

	"jobs-crawler/internal/api"
	"jobs-crawler/internal/config"
	"jobs-crawler/internal/service"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
	"jobs-crawler/internal/sources/lever"
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
		},
		Lever: []lever.Client{
			{Company: "quintoandar"},
			{Company: "loggi"},
		},
	}

	jobs, err := crawler.Run()
	if err != nil {
		log.Printf("crawler error: %v", err)
	}

	fmt.Printf("\n✅ Total QA jobs: %d\n", len(jobs))

	r := api.NewRouter(jobs)
	r.Run(":8080")
}

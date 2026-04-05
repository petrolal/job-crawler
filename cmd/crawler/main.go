package main

import (
	"fmt"

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

	jobs, _ := crawler.Run()

	for _, j := range jobs {
		fmt.Printf(
			"\n🧪 %s\n🌍 Remoto:%v | BSB:%v\n🔗 %s\n",
			j.Title, j.IsRemote, j.IsHybridBrasilia, j.URL,
		)
	}

	fmt.Printf("\n✅ Total final: %d\n", len(jobs))
}

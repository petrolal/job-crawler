package main

import (
	"fmt"

	"jobs-crawler/internal/config"
	"jobs-crawler/internal/service"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
	"jobs-crawler/internal/sources/lever"

	"github.com/gin-gonic/gin"
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

	fmt.Printf("\n✅ Total QA jobs: %d\n", len(jobs))

	g := gin.Default()

	g.GET("/jobs", func(ctx *gin.Context) {
		ctx.JSON(200, jobs)
	})

	g.Run(":8080")
}

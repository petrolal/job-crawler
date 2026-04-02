package main

import (
	"fmt"

	"jobs-crawler/internal/service"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
)

func main() {
	crawler := service.CrawlerService{
		Adzuna: adzuna.Client{
			AppID:  "SEU_ID",
			APIKey: "SUA_KEY",
		},
		Greenhouse: []greenhouse.Client{
			{Company: "nubank"},
			{Company: "vtex"},
		},
	}

	jobs, _ := crawler.Run()

	for _, j := range jobs {
		fmt.Printf(
			"\n🧪 %s\n⭐ %d\n🌍 Remoto:%v | BSB:%v\n🔗 %s\n",
			j.Title, j.Score, j.IsRemote, j.IsHybridBrasilia, j.URL,
		)
	}
}

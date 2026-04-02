package service

import (
	"fmt"
	"sort"

	"jobs-crawler/internal/dedup"
	"jobs-crawler/internal/domain"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
)

type CrawlerService struct {
	Adzuna     adzuna.Client
	Greenhouse []greenhouse.Client
}

func (s CrawlerService) Run() ([]domain.Job, error) {
	var jobs []domain.Job

	// =====================
	// ADZUNA
	// =====================
	adzJobs, err := s.Adzuna.Fetch()
	if err != nil {
		return nil, err
	}

	for _, r := range adzJobs {
		jobs = append(jobs, adzuna.MapToDomain(r))
	}
	fmt.Println("📦 Adzuna:", len(adzJobs))

	// =====================
	// GREENHOUSE
	// =====================
	for _, gh := range s.Greenhouse {
		ghJobs, err := gh.Fetch()
		if err != nil {
			continue
		}
		for _, r := range ghJobs {
			jobs = append(jobs, greenhouse.MapToDomain(gh.Company, r))
		}
	}

	fmt.Println("📦 Total bruto:", len(jobs))

	// =====================
	// DEDUPLICAÇÃO
	// =====================
	jobs = dedup.Deduplicate(jobs)

	// =====================
	// FILTRO QA (CORRETO)
	// =====================
	var qaJobs []domain.Job
	for _, j := range jobs {
		if j.IsLikelyQA {
			qaJobs = append(qaJobs, j)
		}
	}

	// =====================
	// ORDENA POR SCORE
	// =====================
	sort.Slice(qaJobs, func(i, j int) bool {
		return qaJobs[i].Score > qaJobs[j].Score
	})

	fmt.Println("✅ QA jobs finais:", len(qaJobs))
	return qaJobs, nil
}

package service

import (
	"jobs-crawler/internal/domain"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
	"jobs-crawler/internal/sources/lever"
)

type CrawlerService struct {
	Adzuna     adzuna.Client
	Greenhouse []greenhouse.Client
	Lever      []lever.Client
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

	// ===================
	// Lever
	// ===================
	for _, lv := range s.Lever {
		lvJobs, err := lv.Fetch()
		if err != nil {
			continue
		}
		for _, r := range lvJobs {
			jobs = append(jobs, lever.MapToDomain(lv.Company, r))
		}
	}

	// =====================
	// Filter QA
	// =====================
	var qaJobs []domain.Job
	for _, j := range jobs {
		if j.IsLikelyQA {
			qaJobs = append(qaJobs, j)
		}
	}

	return qaJobs, nil
}

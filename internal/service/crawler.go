package service

import (
	"sync"

	"jobs-crawler/internal/domain"
	"jobs-crawler/internal/sources/adzuna"
	"jobs-crawler/internal/sources/greenhouse"
	"jobs-crawler/internal/sources/lever"
)

type CrawlerService struct {
	Adzuna     adzuna.Client
	Greenhouse []greenhouse.Client
	Lever      []lever.Client
	Jobs       []domain.Job
}

func (s CrawlerService) Run() ([]domain.Job, error) {
	var jobs []domain.Job
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)

	go func() {
		defer wg.Done()
		s.getAdzunaJobs(&jobs, &mu)
	}()

	go func() {
		defer wg.Done()
		s.getGreenhouseJobs(&jobs, &mu)
	}()

	go func() {
		defer wg.Done()
		s.getLeverJobs(&jobs, &mu)
	}()

	wg.Wait()

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

// =====================
// ADZUNA
// =====================
func (s *CrawlerService) getAdzunaJobs(jobs *[]domain.Job, mu *sync.Mutex) {
	adzJobs, err := s.Adzuna.Fetch()
	if err != nil {
		return
	}

	for _, r := range adzJobs {
		mu.Lock()
		*jobs = append(*jobs, adzuna.MapToDomain(r))
		mu.Unlock()
	}
}

// =====================
// GREENHOUSE
// =====================
func (s *CrawlerService) getGreenhouseJobs(jobs *[]domain.Job, mu *sync.Mutex) {
	for _, gh := range s.Greenhouse {
		ghJobs, err := gh.Fetch()
		if err != nil {
			continue
		}

		for _, r := range ghJobs {
			mu.Lock()
			*jobs = append(*jobs, greenhouse.MapToDomain(gh.Company, r))
			mu.Unlock()
		}
	}
}

// ===================
// Lever
// ===================

func (s *CrawlerService) getLeverJobs(jobs *[]domain.Job, mu *sync.Mutex) {
	for _, lv := range s.Lever {
		lvJobs, err := lv.Fetch()
		if err != nil {
			continue
		}

		for _, r := range lvJobs {
			mu.Lock()
			*jobs = append(*jobs, lever.MapToDomain(lv.Company, r))
			mu.Unlock()
		}
	}
}

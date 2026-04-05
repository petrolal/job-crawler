package service

import (
	"log"
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
}

func (s CrawlerService) Run() ([]domain.Job, error) {
	var all []domain.Job
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)
	go func() { defer wg.Done(); s.fetchAdzuna(&all, &mu) }()
	go func() { defer wg.Done(); s.fetchGreenhouse(&all, &mu) }()
	go func() { defer wg.Done(); s.fetchLever(&all, &mu) }()
	wg.Wait()

	var jobs []domain.Job
	for _, j := range all {
		if j.IsBrazil || j.IsRemote {
			jobs = append(jobs, j)
		}
	}
	return jobs, nil
}

func (s *CrawlerService) fetchAdzuna(jobs *[]domain.Job, mu *sync.Mutex) {
	results, err := s.Adzuna.Fetch()
	if err != nil {
		return
	}
	for _, r := range results {
		mu.Lock()
		*jobs = append(*jobs, adzuna.MapToDomain(r))
		mu.Unlock()
	}
}

func (s *CrawlerService) fetchGreenhouse(jobs *[]domain.Job, mu *sync.Mutex) {
	for _, gh := range s.Greenhouse {
		results, err := gh.Fetch()
		if err != nil {
			log.Printf("greenhouse [%s]: %v", gh.Company, err)
			continue
		}
		for _, r := range results {
			mu.Lock()
			*jobs = append(*jobs, greenhouse.MapToDomain(gh.Company, r))
			mu.Unlock()
		}
	}
}

func (s *CrawlerService) fetchLever(jobs *[]domain.Job, mu *sync.Mutex) {
	for _, lv := range s.Lever {
		results, err := lv.Fetch()
		if err != nil {
			log.Printf("lever [%s]: %v", lv.Company, err)
			continue
		}
		for _, r := range results {
			mu.Lock()
			*jobs = append(*jobs, lever.MapToDomain(lv.Company, r))
			mu.Unlock()
		}
	}
}

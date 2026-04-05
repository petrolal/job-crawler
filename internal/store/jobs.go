package store

import (
	"sync"

	"jobs-crawler/internal/domain"
)

type CrawlStatus string

const (
	StatusCrawling CrawlStatus = "crawling"
	StatusDone     CrawlStatus = "done"
)

type JobStore struct {
	mu     sync.RWMutex
	jobs   []domain.Job
	status CrawlStatus
}

func NewJobStore() *JobStore {
	return &JobStore{status: StatusCrawling}
}

func (s *JobStore) SetJobs(jobs []domain.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs = jobs
	s.status = StatusDone
}

func (s *JobStore) Get() ([]domain.Job, CrawlStatus) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.jobs, s.status
}

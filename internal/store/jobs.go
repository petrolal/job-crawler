// Package store provides a thread-safe in-memory store for job postings
// collected by the crawler.
package store

import (
	"sync"

	"jobs-crawler/internal/domain"
)

// CrawlStatus represents the current state of the crawl operation.
type CrawlStatus string

const (
	// StatusCrawling indicates the crawler is still fetching jobs.
	StatusCrawling CrawlStatus = "crawling"
	// StatusDone indicates the crawler has finished and results are ready.
	StatusDone CrawlStatus = "done"
)

// JobStore holds the crawled jobs and the current crawl status.
type JobStore struct {
	mu     sync.RWMutex
	jobs   []domain.Job
	status CrawlStatus
}

// NewJobStore returns a new JobStore with status set to crawling.
func NewJobStore() *JobStore {
	return &JobStore{status: StatusCrawling}
}

// SetJobs replaces the stored jobs and marks the crawl as done.
func (s *JobStore) SetJobs(jobs []domain.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs = jobs
	s.status = StatusDone
}

// Get returns the current list of jobs and the crawl status.
func (s *JobStore) Get() ([]domain.Job, CrawlStatus) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.jobs, s.status
}

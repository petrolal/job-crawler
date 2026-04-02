package storage

import "jobs-crawler/internal/domain"

type MemoryStore struct {
	Jobs []domain.Job
}

func (m *MemoryStore) Save(j domain.Job) {
	m.Jobs = append(m.Jobs, j)
}

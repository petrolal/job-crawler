package adzuna

import (
	"jobs-crawler/internal/classifier"
	"jobs-crawler/internal/domain"
)

func MapToDomain(r AdzunaJob) domain.Job {
	text := r.Title + " " + r.Description

	job := domain.Job{
		ID:          r.ID,
		Title:       r.Title,
		Company:     r.Company.DisplayName,
		Location:    r.Location.DisplayName,
		Description: r.Description,
		URL:         r.RedirectURL,
		Source:      "adzuna",
	}

	job.IsRemote = classifier.IsRemote(text)
	job.IsHybridBrasilia = classifier.IsHybridBrasilia(
		r.Title, r.Description, r.Location.DisplayName)

	job.IsLikelyQA = classifier.IsLikelyQA(
		job.Title, job.Description)

	return job
}

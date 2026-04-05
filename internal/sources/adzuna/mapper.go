package adzuna

import (
	"jobs-crawler/internal/classifier"
	"jobs-crawler/internal/domain"
)

func MapToDomain(r AdzunaJob) domain.Job {
	text := r.Title + " " + r.Description
	return domain.Job{
		ID:          r.ID,
		Title:       r.Title,
		Company:     r.Company.DisplayName,
		Location:    r.Location.DisplayName,
		Description: r.Description,
		URL:         r.RedirectURL,
		Source:      "adzuna",
		IsRemote:    classifier.IsRemote(text),
		IsBrazil:    classifier.IsBrazil(r.Title, r.Description, r.Location.DisplayName),
	}
}

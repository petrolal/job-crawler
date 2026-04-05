package lever

import (
	"jobs-crawler/internal/classifier"
	"jobs-crawler/internal/domain"
)

func MapToDomain(company string, r LeverJob) domain.Job {
	fullText := r.Text + " " + r.Description

	return domain.Job{
		ID:          r.ID,
		Title:       r.Text,
		Company:     company,
		Location:    r.Categories.Location,
		Description: r.Description,
		URL:         r.HostedURL,
		Source:      "lever",

		IsRemote: classifier.IsRemote(fullText),
		IsBrazil: classifier.IsBrazil(
			r.Text,
			r.Description,
			r.Categories.Location,
		),
		IsLikelyQA: classifier.IsLikelyQA(r.Text, r.Description),
	}
}

package lever

import (
	"jobs-crawler/internal/classifier"
	"jobs-crawler/internal/domain"
)

func MapToDomain(company string, r LeverJob) domain.Job {
	text := r.Text + " " + r.Description
	return domain.Job{
		ID:          r.ID,
		Title:       r.Text,
		Company:     company,
		Location:    r.Categories.Location,
		Description: r.Description,
		URL:         r.HostedURL,
		Source:      "lever",
		IsRemote:    classifier.IsRemote(text),
		IsBrazil:    classifier.IsBrazil(r.Text, r.Description, r.Categories.Location),
	}
}

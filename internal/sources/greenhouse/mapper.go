package greenhouse

import (
	"strconv"

	"jobs-crawler/internal/classifier"
	"jobs-crawler/internal/domain"
)

func MapToDomain(company string, r GHJob) domain.Job {
	text := r.Title + " " + r.Content

	job := domain.Job{
		ID:          company + "-" + strconv.Itoa(r.ID),
		Title:       r.Title,
		Company:     company,
		Location:    r.Location.Name,
		Description: r.Content,
		URL:         r.AbsoluteURL,
		Source:      "greenhouse",
	}

	job.IsRemote = classifier.IsRemote(text)
	job.IsBrazil = classifier.IsBrazil(
		r.Title, r.Content, r.Location.Name)

	job.IsLikelyQA = classifier.IsLikelyQA(
		job.Title, job.Description)

	return job
}

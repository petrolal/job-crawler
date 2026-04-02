package dedup

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"jobs-crawler/internal/domain"
)

func Deduplicate(jobs []domain.Job) []domain.Job {
	seen := map[string]domain.Job{}

	for _, job := range jobs {
		key := hash(job.Company + normalize(job.Title))

		existing, ok := seen[key]
		if !ok || job.Score > existing.Score {
			seen[key] = job
		}
	}

	var out []domain.Job
	for _, j := range seen {
		out = append(out, j)
	}
	return out
}

func normalize(t string) string {
	r := strings.NewReplacer(
		"senior", "", "sr", "", "jr", "", "pleno", "",
	)
	return r.Replace(strings.ToLower(t))
}

func hash(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

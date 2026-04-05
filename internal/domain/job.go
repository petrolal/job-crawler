// Package domain defines the core data types used across the application.
package domain

// Job represents a job posting collected from an external source.
type Job struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Source      string `json:"source"`

	IsRemote bool `json:"-"`
	IsBrazil bool `json:"-"`
}

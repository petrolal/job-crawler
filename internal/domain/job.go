package domain

type Job struct {
	ID          string
	Title       string
	Company     string
	Location    string
	Description string
	URL         string
	Source      string

	IsRemote   bool
	IsBrazil   bool
	IsLikelyQA bool
}

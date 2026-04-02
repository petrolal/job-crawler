package adzuna

type AdzunaResponse struct {
	Results []AdzunaJob `json:"results"`
}

type AdzunaJob struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	RedirectURL string `json:"redirect_url"`
	Company     struct {
		DisplayName string `json:"display_name"`
	} `json:"company"`
	Location struct {
		DisplayName string `json:"display_name"`
	} `json:"location"`
}

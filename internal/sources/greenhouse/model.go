package greenhouse

type GHResponse struct {
	Jobs []GHJob `json:"jobs"`
}

type GHJob struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Content     string `json:"content"`
	AbsoluteURL string `json:"absolute_url"`
}

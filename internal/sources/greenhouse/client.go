package greenhouse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Company string
}

func (c Client) Fetch() ([]GHJob, error) {
	url := fmt.Sprintf(
		"https://boards-api.greenhouse.io/v1/boards/%s/jobs",
		c.Company,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r GHResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r.Jobs, err
}

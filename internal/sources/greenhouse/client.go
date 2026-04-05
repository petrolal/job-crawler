package greenhouse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Company string
}

func (c Client) Fetch() ([]GHJob, error) {
	url := fmt.Sprintf(
		"https://boards-api.greenhouse.io/v1/boards/%s/jobs?content=true",
		c.Company,
	)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r GHResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r.Jobs, err
}

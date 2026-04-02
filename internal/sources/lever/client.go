package lever

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Company string
}

func (c Client) Fetch() ([]LeverJob, error) {
	url := fmt.Sprintf(
		"https://api.lever.co/v0/postings/%s?mode=json",
		c.Company,
	)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jobs []LeverJob
	if err := json.NewDecoder(resp.Body).Decode(&jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}

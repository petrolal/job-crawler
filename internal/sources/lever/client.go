package lever

import (
	"encoding/json"
	"fmt"
	"io"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Lever returns {"ok":false,"error":"..."} when the company slug is invalid
	if len(body) > 0 && body[0] == '{' {
		return nil, fmt.Errorf("lever: unexpected response for company %q (check the slug)", c.Company)
	}

	var jobs []LeverJob
	if err := json.Unmarshal(body, &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}

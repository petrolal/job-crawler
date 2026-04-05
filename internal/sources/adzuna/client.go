package adzuna

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	AppID       string
	APIKey      string
	ResultsPage int
	Pages       int
}

func (c Client) Fetch() ([]AdzunaJob, error) {
	var all []AdzunaJob

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	for page := 1; page <= c.Pages; page++ {
		u, err := url.Parse(
			fmt.Sprintf(
				"https://api.adzuna.com/v1/api/jobs/br/search/%d",
				page,
			),
		)
		if err != nil {
			return nil, err
		}

		// ✅ Query params (como você fazia antes)
		q := u.Query()
		q.Set("what", "qa")
		q.Set("app_id", c.AppID)
		q.Set("app_key", c.APIKey)
		q.Set("results_per_page", strconv.Itoa(c.ResultsPage))

		u.RawQuery = q.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Accept", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		var response AdzunaResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			resp.Body.Close()
			return nil, err
		}

		resp.Body.Close()

		if len(response.Results) == 0 {
			fmt.Println("no Results on Azuna")
			break // ✅ sem mais resultados
		}

		all = append(all, response.Results...)

		// ✅ throttle leve
		time.Sleep(600 * time.Millisecond)
	}

	return all, nil
}

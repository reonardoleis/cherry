package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Country struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Population int     `json:"population"`
	LandArea   int     `json:"land_area"`
	Density    float64 `json:"density"`
	Capital    string  `json:"capital"`
	Currency   string  `json:"currency"`
	Flag       string  `json:"flag"`
}

func FetchCountry(id int) (*Country, error) {
	resp, err := http.Get(fmt.Sprintf("https://freetestapi.com/api/v1/countries/%d", id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var country Country
	err = json.Unmarshal(body, &country)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &country, nil
}

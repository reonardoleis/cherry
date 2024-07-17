package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Country struct {
	Capital string `json:"capital"`
}

func searchCountry(query string) (*Country, error) {
	url := fmt.Sprintf("https://freetestapi.com/api/v1/countries?search=%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var countries []Country
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}

	if len(countries) > 0 {
		return &countries[0], nil
	}

	return nil, fmt.Errorf("no countries found")
}

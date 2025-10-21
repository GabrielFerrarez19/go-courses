package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Results struct {
	Search       []SearchResults `json:"Search"`
	TotalResults string          `json:"totalResults"`
	Response     string          `json:"Response"`
}
type SearchResults struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Search(apiKey, title string) (Results, error) {
	v := url.Values{}

	v.Set("apikey", apiKey)
	v.Set("s", title)

	resp, err := http.Get("http://www.omdbapi.com/?" + v.Encode())
	if err != nil {
		return Results{}, fmt.Errorf("failed to make request to omdb: %w", err)
	}
	defer resp.Body.Close()

	var result Results

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Results{}, fmt.Errorf("failed to decode response from omdb %w", err)
	}

	return result, nil
}

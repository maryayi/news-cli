package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	everythingURL  = "https://newsapi.org/v2/everything"
	topHeadlineURL = "https://newsapi.org/v2/top-headlines"
)

func FetchHeadlines(apiKey, keyword string, count int) ([]Article, error) {
	params := url.Values{}
	params.Set("pageSize", fmt.Sprintf("%d", count))
	params.Set("language", "en")
	params.Set("apiKey", apiKey)

	var endpoint string
	if keyword == "" {
		endpoint = topHeadlineURL
	} else {
		endpoint = everythingURL
		params.Set("q", keyword)
		params.Set("sortBy", "publishedAt")
	}

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var result APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if result.Status != "ok" {
		return nil, fmt.Errorf("API error: %s", result.Message)
	}

	return result.Articles, nil
}

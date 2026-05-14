package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const baseURL = "https://newsapi.org/v2/everything"

func FetchHeadlines(apiKey, keyword string, count int) ([]Article, error) {
	params := url.Values{}
	params.Set("q", keyword)
	params.Set("pageSize", fmt.Sprintf("%d", count))
	params.Set("sortBy", "publishedAt")
	params.Set("language", "en")
	params.Set("apiKey", apiKey)

	resp, err := http.Get(baseURL + "?" + params.Encode())
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

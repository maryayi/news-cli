package news

type Source struct {
	Name string `json:"name"`
}

type Article struct {
	Source      Source `json:"source"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
}

type APIResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
	Message      string    `json:"message"`
}

package model

type URL struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Alias  string `json:"alias"`
	Clicks int    `json:"clicks"`
}

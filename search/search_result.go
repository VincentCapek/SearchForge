package search

type SearchResult struct {
	URL string `json:"url"`
	Title string `json:"title"`
	Domain string `json:"domain"`
	Snippet string `json:"snippet"`
}

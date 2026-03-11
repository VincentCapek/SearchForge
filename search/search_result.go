package search

type SearchResultDuckDuckGo struct {
	URL string `json:"url"`
	Title string `json:"title"`
	Domain string `json:"domain"`
	Snippet string `json:"snippet"`
}

type SearchResultBrave struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Domain   string `json:"domain"`
	SiteName string `json:"siteName,omitempty"`
	Author   string `json:"author,omitempty"`
}

type SearchResultQwant struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Domain  string `json:"domain"`
	Snippet string `json:"snippet"`
}
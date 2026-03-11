package search

import (
	"fmt"
	"net/url"
	"strings"

	"serpgo/utils"

	"github.com/PuerkitoBio/goquery"
)

const ddgSearchEndpoint = "https://html.duckduckgo.com/html/?q="

func GetDDGResults(query string) ([]SearchResultDuckDuckGo, error) {
	searchURL := ddgSearchEndpoint + url.QueryEscape(query)

	doc, err := utils.FetchDocument(searchURL)
	if err != nil {
		return nil, err
	}

	results := make([]SearchResultDuckDuckGo, 0)

	doc.Find(".result").Each(func(i int, item *goquery.Selection) {
		if item.HasClass("result--ad") {
			return
		}

		titleEl := item.Find(".result__title a")
		title := strings.TrimSpace(titleEl.Text())
		link := utils.NormalizeDuckDuckGoURL(titleEl.AttrOr("href", ""))
		snippet := strings.TrimSpace(item.Find(".result__snippet").Text())

		if title == "" || link == "" {
			return
		}

		parsedURL, err := url.Parse(link)
		if err != nil {
			return
		}

		host := parsedURL.Hostname()
		if host == "" || isBlockedSite(host) {
			return
		}

		results = append(results, SearchResultDuckDuckGo{
			Title: title,
			URL: link,
			Domain: host,
			Snippet: snippet,
		})
	})

	if len(results) == 0 {
		return nil, fmt.Errorf("aucun résultat trouvé ou page bloquée")
	}

	return results, nil
}
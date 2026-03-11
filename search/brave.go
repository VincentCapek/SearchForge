package search

import (
	"fmt"
	"net/url"
	"strings"

	"serpgo/utils"

	"github.com/PuerkitoBio/goquery"
)

func GetBraveResults(q string) ([]SearchResultBrave, error) {
	searchURL := fmt.Sprintf("https://search.brave.com/search?q=%s", url.QueryEscape(q))
	doc, err := utils.FetchDocument(searchURL)
	if err != nil {
		return nil, err
	}

	results := make([]SearchResultBrave, 0)
	seen := make(map[string]struct{})

	doc.Find("div.snippet:not(.standalone)").Each(func(_ int, item *goquery.Selection) {
		a := item.Find("a").First()
		link, _ := a.Attr("href")
		link, _ = url.QueryUnescape(link)

		if link == "" || link == "#" || strings.HasPrefix(link, "/") {
			return
		}

		u, err := url.Parse(link)
		if err != nil {
			return
		}

		if _, ok := seen[link]; ok {
			return
		}

		if isBlockedSite(u.Hostname()) {
			return
		}

		seen[link] = struct{}{}	

		results = append(results, SearchResultBrave{
			URL:      link,
			Title:    strings.TrimSpace(item.Find(".title").First().Text()),
			Desc:     strings.TrimSpace(item.Find(".snippet-description").Text()),
			Domain:   u.Host,
			SiteName: strings.TrimSpace(item.Find(".netloc").Text()),
			Author:   strings.TrimSpace(item.Find(".snippet-url").First().Text()),
		})
	})

	return results, nil
}
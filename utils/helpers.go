package utils

import (
	"net/url"
	"strings"
)

func NormalizeDuckDuckGoURL(raw string) string {
	if raw == "" {
		return ""
	}

	if strings.HasPrefix(raw, "//") {
		raw = "https:" + raw
	}

	parsedURL, err := url.Parse(raw)
	if err != nil {
		return raw
	}

	if uddg := parsedURL.Query().Get("uddg"); uddg != "" {
		decoded, err := url.QueryUnescape(uddg)
		if err == nil {
			return decoded
		}
	}

	return raw
}
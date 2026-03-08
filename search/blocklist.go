package search

import "strings"

var blocklist = []string{
	"pinterest.com",
	"pinterest.fr",
	"allocine.com",
	"jeuxvideo.com",
	"w3schools.com",
}

func isBlockedSite(host string) bool {
	host = strings.ToLower(strings.TrimPrefix(host, "www."))

	for _, blocked := range blocklist {
		blocked = strings.ToLower(strings.TrimSuffix(blocked, "www."))

		if host == blocked || strings.HasSuffix(host, "."+blocked) {
			return true
		}
	}
	return false
}
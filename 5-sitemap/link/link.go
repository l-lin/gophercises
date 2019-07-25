package link

import (
	"net/url"
	"strings"
)

// Link represents the content of a hypermedia link from a HTML file
type Link struct {
	Href url.URL
}

// New instanciates a new link
func New(rawurl string) (*Link, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return &Link{Href: *u}, nil
}

func (l *Link) String() string {
	return l.Href.String()
}

// HasSameDomain checks if the link has the same domain as the given URL
func (l *Link) HasSameDomain(u url.URL) bool {
	return l.Href.Host == u.Host
}

// IsSameLink checks if the given link is the same URL
func (l *Link) IsSameLink(target Link) bool {
	return strings.Trim(l.Href.String(), " ") == strings.Trim(target.Href.String(), " ")
}

// Unique returns a set of links without duplicate links
func Unique(links []Link) []Link {
	keys := make(map[string]bool)
	result := []Link{}
	for _, l := range links {
		if _, value := keys[l.Href.String()]; !value {
			keys[l.Href.String()] = true
			result = append(result, l)
		}
	}
	return result
}

package sitemap

import (
	"encoding/xml"

	"github.com/l-lin/5-sitemap/link"
)

const sitemapSchema = "http://www.sitemaps.org/schemas/sitemap/0.9"

// Sitemap represents the map of all pages within a specific domain
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

// URL represents a URL in the sitemap
type URL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

// New sitemap
func New() Sitemap {
	return Sitemap{Xmlns: sitemapSchema, URLs: []URL{}}
}

// FromLinks generate a sitemap from a given links
func FromLinks(links []link.Link) Sitemap {
	s := New()
	for _, l := range links {
		u := URL{Loc: l.Href.String()}
		s.URLs = append(s.URLs, u)
	}
	return s
}

// ToXML transforms the sitemap into a XML format
func (s Sitemap) ToXML() ([]byte, error) {
	result, err := xml.MarshalIndent(s, "", "    ")
	if err != nil {
		return nil, err
	}
	result = []byte(xml.Header + string(result))
	return result, nil
}

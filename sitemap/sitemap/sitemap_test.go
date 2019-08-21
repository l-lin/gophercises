package sitemap

import (
	"fmt"
	"testing"

	"github.com/l-lin/sitemap/link"
)

func TestSitemap(t *testing.T) {
	links := []link.Link{}
	l, _ := link.New("https://httpbin.org")
	links = append(links, *l)
	l, _ = link.New("http://foo.com/bar")
	links = append(links, *l)

	s := FromLinks(links)
	result, err := s.ToXML()
	if err != nil {
		t.Errorf("%v", err)
	}
	expected := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://httpbin.org</loc>
    </url>
    <url>
        <loc>http://foo.com/bar</loc>
    </url>
</urlset>`
	if expected != fmt.Sprintf("%s", result) {
		t.Errorf("expected:\n%s\nactual:\n%s\n", expected, result)
	}
}

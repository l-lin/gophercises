package link

import (
	"io"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

// Parse links from given reader
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := parseLinks(doc)
	return links, nil
}

func parseLinks(n *html.Node) []Link {
	links := make([]Link, 0)
	if isLink(n) {
		href := fetchAttr(n, "href")
		l, err := New(href)
		if err != nil {
			logrus.Warnln(err)
		} else {
			links = append(links, *l)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseLinks(c)...)
	}
	return links
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func fetchAttr(n *html.Node, attrKey string) string {
	for _, a := range n.Attr {
		if a.Key == attrKey {
			return a.Val
		}
	}
	return ""
}

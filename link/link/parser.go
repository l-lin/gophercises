package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Parse file to fetch links
func Parse(r io.Reader) ([]*Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := parseLinks(doc)
	return links, nil
}

func parseLinks(n *html.Node) []*Link {
	links := make([]*Link, 0)
	if isLink(n) {
		l := &Link{}
		l.Href = fetchAttr(n, "href")
		l.Texts = append(l.Texts, fetchText(n)...)
		links = append(links, l)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseLinks(c)...)
	}
	return links
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func fetchText(n *html.Node) []string {
	texts := make([]string, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isText(c) && strings.TrimSpace(c.Data) != "" {
			texts = append(texts, c.Data)
		} else {
			texts = append(texts, fetchText(c)...)
		}
	}
	return texts
}

func isText(n *html.Node) bool {
	return n.Type == html.TextNode
}

func fetchAttr(n *html.Node, attrKey string) string {
	for _, a := range n.Attr {
		if a.Key == attrKey {
			return a.Val
		}
	}
	return ""
}

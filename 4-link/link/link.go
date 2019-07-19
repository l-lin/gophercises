package link

import (
	"fmt"
	"strings"
)

// Link represents the content of a hypermedia link from a HTML file
type Link struct {
	Href  string
	Texts []string
}

// GetText by concatenating all texts from link child nodes
func (l *Link) GetText() string {
	if len(l.Texts) == 0 {
		return ""
	}

	var text strings.Builder
	for i := 0; i < len(l.Texts)-1; i++ {
		t := l.Texts[i]
		text.WriteString(strings.TrimSpace(t))
		text.WriteString(" ")
	}
	t := l.Texts[len(l.Texts)-1]
	text.WriteString(strings.TrimSpace(t))
	return text.String()
}

func (l *Link) String() string {
	return fmt.Sprintf(`Link{
  Href: "%s",
  Text: "%s",
}`, l.Href, l.GetText())
}

func newLink() *Link {
	return &Link{
		Texts: make([]string, 0),
	}
}

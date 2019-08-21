package story

import (
	"encoding/json"
	"io/ioutil"
)

// Chapter of a story
type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []*Option `json:"options"`
}

// Option to choose at the end of the chapter
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// ReadFromFile parses and transforms into a slice of Chapters
func ReadFromFile(inputFile string) (map[string]*Chapter, error) {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	var story map[string]*Chapter
	err = json.Unmarshal(content, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

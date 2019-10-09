// Package hn implements a really basic Hacker News client
package hn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiBase = "https://hacker-news.firebaseio.com/v0"
)

// Client is an API client used to interact with the Hacker News API
type Client struct {
	// unexported fields...
	apiBase string
}

// Making the Client zero value useful without forcing users to do something
// like `NewClient()`
func (c *Client) defaultify() {
	if c.apiBase == "" {
		c.apiBase = apiBase
	}
}

// TopItems returns the ids of roughly 450 top items in decreasing order. These
// should map directly to the top 450 things you would see on HN if you visited
// their site and kept going to the next page.
//
// TopItems does not filter out job listings or anything else, as the type of
// each item is unknown without further API calls.
func (c *Client) TopItems() ([]int, error) {
	c.defaultify()
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apiBase))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ids []int
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// GetItem will return the Item defined by the provided ID.
func (c *Client) GetItem(id int) (Item, error) {
	c.defaultify()
	var item Item
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiBase, id))
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// Item represents a single item returned by the HN API. This can have a type
// of "story", "comment", or "job" (and probably more values), and one of the
// URL or Text fields will be set, but not both.
//
// For the purpose of this exercise, we only care about items where the
// type is "story", and the URL is set.
type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`

	// Only one of these should exist
	Text string `json:"text"`
	URL  string `json:"url"`
	// specific attribute
	Host string
}

// GetTopStories fetches the HN top stories
func GetTopStories(numStories int, timeout time.Duration) ([]Item, error) {
	var client Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, err
	}
	type result struct {
		i Item
		j int
	}
	t := time.After(timeout)
	computedNbStories := computeNbStoriesToFetch(numStories)
	items := make([]Item, computedNbStories)
	resultCh := make(chan result, computedNbStories)
	for j := 0; j < computedNbStories; j++ {
		id := ids[j]
		go func(id, j int, resultCh chan result) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				return
			}
			i := parseHNItem(hnItem)
			resultCh <- result{i, j}
		}(id, j, resultCh)
	}
	nbStoriesFetched := 0
	for nbStoriesFetched < computedNbStories {
		select {
		case r := <-resultCh:
			nbStoriesFetched++
			items[r.j] = r.i
		case <-t:
			log.Printf("%v timeout...\n", timeout)
			break
		}
	}
	stories := []Item{}
	for _, i := range items {
		if len(stories) >= numStories {
			break
		}
		if isStoryLink(i) {
			stories = append(stories, i)
		}
	}
	return stories, nil
}

func computeNbStoriesToFetch(numStories int) int {
	return int(float64(numStories) * 1.25)
}

func isStoryLink(item Item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(item Item) Item {
	url, err := url.Parse(item.URL)
	if err == nil {
		item.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return item
}

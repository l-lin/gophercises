package crawler

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/l-lin/gophercises/sitemap/link"
	log "github.com/sirupsen/logrus"
)

var timer = 30 * time.Second

// Crawler crawls the site
type Crawler struct {
	// URL is the given URL
	URL *url.URL
}

// New instanciates a new crawler with an absolute URL given in parameter
func New(rawurl string) (*Crawler, error) {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if !u.IsAbs() {
		return nil, fmt.Errorf("The given URL %s must be absolute", rawurl)
	}

	return &Crawler{URL: u}, nil
}

// Perform site crawling to get links
func (c *Crawler) Perform(depth int) ([]link.Link, error) {
	crawledURLs := make(map[string]bool)
	urls := []url.URL{*c.URL}
	links := make([]link.Link, 0)
	i := 0
	for i < depth {
		log.WithField("depth", i).Info("Crawling iteration")
		// Do not crawl already crawled URLs
		urls = filterNotCrawledURLs(urls, crawledURLs)
		if len(urls) == 0 {
			break
		}
		crawledLinks, err := crawlURLs(urls)
		if err != nil {
			return nil, err
		}
		crawledLinks = link.Unique(crawledLinks)
		// Mark URLs as crawled
		for _, u := range urls {
			crawledURLs[u.String()] = true
		}
		links = append(links, crawledLinks...)
		urls = []url.URL{}
		for _, crawledLink := range crawledLinks {
			urls = append(urls, crawledLink.Href)
		}
		i++
	}
	return link.Unique(links), nil
}

func filterNotCrawledURLs(urls []url.URL, crawledURLs map[string]bool) []url.URL {
	result := []url.URL{}
	for _, u := range urls {
		if _, ok := crawledURLs[u.String()]; !ok {
			result = append(result, u)
		}
	}
	return result
}

func crawlURLs(urls []url.URL) ([]link.Link, error) {
	ch := make(chan []link.Link, len(urls))
	errCh := make(chan error, 1)
	for _, u := range urls {
		go crawl(ch, errCh, u)
	}
	links := make([]link.Link, 0)
	select {
	case l := <-ch:
		links = append(links, l...)
		break
	case err := <-errCh:
		return nil, err
	case <-time.After(timer):
		return nil, errors.New("Timed out")
	}
	return links, nil
}

func crawl(ch chan<- []link.Link, errCh chan<- error, u url.URL) {
	log.WithField("url", u.String()).Info("Fetching links from URL...")
	resp, err := http.Get(u.String())
	if err != nil {
		errCh <- err
		return
	}
	if resp.StatusCode != 200 {
		errCh <- fmt.Errorf("Could not fetch content of %s. Error status code was %d", u.RawPath, resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCh <- err
		return
	}
	links, err := link.Parse(bytes.NewReader(body))
	if err != nil {
		errCh <- err
		return
	}
	result := make([]link.Link, 0)
	for _, l := range links {
		if !l.Href.IsAbs() {
			crawledURL, err := url.Parse(u.String())
			if err != nil {
				errCh <- err
				return
			}
			crawledURL.Path = l.Href.Path
			result = append(result, link.Link{Href: *crawledURL})
		}
		if l.HasSameDomain(u) {
			result = append(result, l)
		}
	}
	ch <- result
}

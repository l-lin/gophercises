package cmd

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/l-lin/gophercises/quiethn/hn"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start quiet HN server",
		Run:   runServe,
	}
	port, numStories int
	timeout          = 10 * time.Second
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "the port to start the web server on")
	serveCmd.Flags().IntVarP(&numStories, "num-stories", "n", 30, "the number of top stories to display")
}

func runServe(cmd *cobra.Command, args []string) {
	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := getTopStories()
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}
		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func getTopStories() ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, err
	}
	story := make(chan item)
	stories := []item{}
	t := time.After(timeout)
	for j := 0; j < computeNbStoriesToFetch(numStories); j++ {
		id := ids[j]
		go func(id int, story chan item) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				return
			}
			i := parseHNItem(hnItem)
			story <- i
		}(id, story)
	}
	for len(stories) < numStories {
		select {
		case s := <-story:
			if isStoryLink(s) {
				stories = append(stories, s)
			}
		case <-t:
			log.Printf("%v timeout...\n", timeout)
			return stories, nil
		}
	}
	return stories, nil
}

func computeNbStoriesToFetch(numStories int) int {
	return int(float64(numStories) * 1.25)
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}

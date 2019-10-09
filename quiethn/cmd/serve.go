package cmd

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
		stories, err := hn.GetTopStories(numStories, timeout)
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

type templateData struct {
	Stories []hn.Item
	Time    time.Duration
}

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var input string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "urlshort",
	Short: "URL shortener web application",
	Run:   run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "map", "input URL mapping")
}

func run(cmd *cobra.Command, args []string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	var pathsToURLs map[string]string
	pathsToURLs = mappingFromMap()
	mapHandler := handler(pathsToURLs, mux)
	log.Fatal(http.ListenAndServe(":8080", mapHandler))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func handler(pathsToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mappedPath := pathsToURLs[r.URL.Path]
		if mappedPath != "" {
			log.Printf("Redirecting to %s", mappedPath)
			http.Redirect(w, r, mappedPath, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func mappingFromMap() map[string]string {
	return map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
}

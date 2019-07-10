package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/l-lin/2-urlshort/mapper"
	"github.com/spf13/cobra"
)

const port = ":8080"

var (
	yamlFile string
	jsonFile string
	dbURL    string
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "urlshort",
		Short: "URL shortener web application",
		Run:   run,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&yamlFile, "yaml-file", "mapping.yaml", "input yaml file")
	rootCmd.PersistentFlags().StringVar(&jsonFile, "json-file", "mapping.json", "input json file")
	rootCmd.PersistentFlags().StringVar(&dbURL, "db-url", "postgres://postgres@localhost:5432/urlshort?sslmode=disable", "DB URL")
}

func run(cmd *cobra.Command, args []string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	// Maps
	pathsToURLs := mapper.FromMap()
	h := handler(pathsToURLs, mux)
	// YAML
	pathsToURLs, err := mapper.FromYaml(yamlFile)
	if err != nil {
		log.Printf("Could not map URLs from Yaml file %s. Error was: %v. Skipping...", yamlFile, err)
	} else {
		h = handler(pathsToURLs, h)
	}
	// JSON
	pathsToURLs, err = mapper.FromJSON(jsonFile)
	if err != nil {
		log.Printf("Could not map URLs from JSON file %s. Error was: %v. Skipping...", jsonFile, err)
	} else {
		h = handler(pathsToURLs, h)
	}
	// DB
	pathsToURLs, err = mapper.FromDB(dbURL)
	if err != nil {
		log.Printf("Could not map URLs from database. Error was: %v. Skipping...", err)
	} else {
		h = handler(pathsToURLs, h)
	}
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, h))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func handler(pathsToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	log.Printf("URLs mapped: %v", pathsToURLs)
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

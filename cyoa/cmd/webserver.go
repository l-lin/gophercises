package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/l-lin/cyoa/story"
	"github.com/spf13/cobra"
)

var (
	webserverCmd = &cobra.Command{
		Use:   "webserver",
		Short: "Start the story in a web server",
		Run:   executeWebServer,
	}
	port int
	s    map[string]*story.Chapter
	tpl  *template.Template
)

func init() {
	rootCmd.AddCommand(webserverCmd)
	webserverCmd.LocalFlags().IntVarP(&port, "port", "p", 8080, "web server port")
}

func executeWebServer(cmd *cobra.Command, args []string) {
	var e error
	s, e = story.ReadFromFile(inputFile)
	if e != nil {
		log.Fatalln(e)
	}
	tpl, e = template.ParseFiles("templates/cyao.html")
	if e != nil {
		log.Println(e)
	}

	server := http.NewServeMux()
	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	server.HandleFunc("/", handler)
	log.Printf("Server started on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		log.Printf("Redirecting to /%s", startArc)
		http.Redirect(w, r, fmt.Sprintf("/%s", startArc), http.StatusFound)
		return
	}
	path := sanitizePath(r.URL.Path)
	chapter, ok := s[path]
	if !ok {
		log.Printf("Could not find chapter %s\n", path)
		http.Redirect(w, r, "/public/html/404.html", http.StatusFound)
		return
	}
	data := make(map[string]interface{}, 0)
	data["Chapter"] = chapter
	data["StartChapter"] = &story.Option{
		Arc:  startArc,
		Text: fmt.Sprintf("Go back to %s", startArc),
	}

	tpl.Execute(w, data)
}

func sanitizePath(path string) string {
	return strings.ReplaceAll(path, "/", "")
}

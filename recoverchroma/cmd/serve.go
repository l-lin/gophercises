package cmd

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/l-lin/gophercises/recoverchroma/source"
	"github.com/spf13/cobra"
)

var (
	port     int
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "start web server",
		Run:   runServe,
	}
)

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "web server port")
	rootCmd.AddCommand(serveCmd)
}

func runServe(cmd *cobra.Command, args []string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/", sourceCode)
	log.Printf("starting web server on port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), devMw(mux)))
}

func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, source.RenderStack(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func sourceCode(w http.ResponseWriter, r *http.Request) {
	paths, ok := r.URL.Query()["path"]
	if !ok || len(paths[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing query parameter 'path'")
		return
	}
	lineNbStr, ok := r.URL.Query()["lineNb"]
	if !ok || len(lineNbStr[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing query parameter 'lineNb'")
		return
	}
	lineNb, err := strconv.Atoi(lineNbStr[0])
	if err != nil {
		log.Fatal(err)
	}
	path := paths[0]
	f, err := source.GetFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = f.RenderTo(w, lineNb)
	if err != nil {
		log.Fatal(err)
	}
}

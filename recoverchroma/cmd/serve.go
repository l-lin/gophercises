package cmd

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

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
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, string(stack))
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

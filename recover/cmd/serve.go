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
	h := handler(mux)
	if devMode {
		log.Printf("starting web server on port %d in development mode", port)
	} else {
		log.Printf("starting web server on port %d", port)
	}
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), h))
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	// Write() is called before WriteHeader() => that's why it's a HTTP 200 that is returned
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func handler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error: %v\n%s", r, debug.Stack())
				w.WriteHeader(http.StatusInternalServerError)
				if devMode {
					fmt.Fprintf(w, string(debug.Stack()))
				} else {
					fmt.Fprintf(w, "Something went wrong")
				}
			}
		}()
		// By having our own ResponseWriter, we are not calling w.Write() first, but
		// w.WriteHeader() => we can safely return HTTP 500
		rw := &responseWriter{ResponseWriter: w}
		h.ServeHTTP(rw, req)
		rw.flush()
	}
}

type responseWriter struct {
	http.ResponseWriter
	w      [][]byte
	status int
}

func (r *responseWriter) Write(b []byte) (int, error) {
	r.w = append(r.w, b)
	return len(b), nil
}

func (r *responseWriter) WriteHeader(status int) {
	r.status = status
}

func (r *responseWriter) flush() error {
	if r.status != 0 {
		r.ResponseWriter.WriteHeader(r.status)
	}
	for _, write := range r.w {
		_, err := r.ResponseWriter.Write(write)
		if err != nil {
			return err
		}
	}
	return nil
}

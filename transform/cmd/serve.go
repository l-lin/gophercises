package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l-lin/gophercises/transform/web"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch web server to upload and transform an image",
		Run:   runServe,
	}
	port int
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "the port to start the web server on")
}

func runServe(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()
	r.Methods("POST").Path("/io/images").Handler(web.UploadHandler())
	r.Handle("/", http.FileServer(http.Dir("./public")))
	h := web.DefaultHandler(r)

	// Start the server
	log.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}

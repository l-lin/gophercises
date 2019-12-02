package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/rogpeppe/fastuuid"
	"github.com/spf13/cobra"
)

const (
	maxUploadSize = 100 * 1024 * 1024 // 100Mb
	uploadPath    = "/tmp"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/images", imageHandler())
	mux.HandleFunc("/", defaultHandler())
	h := handler(mux)

	// Start the server
	log.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}

func imageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if http.MethodPost == r.Method {
			r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
			if err := r.ParseMultipartForm(maxUploadSize); err != nil {
				displayError(w, "File too big", http.StatusBadRequest)
				return
			}
			file, _, err := r.FormFile("file")
			if err != nil {
				displayError(w, fmt.Sprintf("Invalid file: %s", err), http.StatusBadRequest)
				return
			}
			defer file.Close()
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				displayError(w, fmt.Sprintf("Invalid file: %s", err), http.StatusBadRequest)
				return
			}
			detectedFileType := http.DetectContentType(fileBytes)
			switch detectedFileType {
			case "image/jpeg", "image/jpg", "image/gif", "image/png":
				break
			default:
				displayError(w, fmt.Sprintf("Invalid file: %s", err), http.StatusBadRequest)
				return
			}
			fileName := fastuuid.MustNewGenerator().Hex128()
			fileEndings, err := mime.ExtensionsByType(detectedFileType)
			if err != nil {
				displayError(w, fmt.Sprintf("Can't read file type: %s", err), http.StatusBadRequest)
				return
			}
			newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
			newFile, err := os.Create(newPath)
			if err != nil {
				displayError(w, fmt.Sprintf("Can't write file type: %s", err), http.StatusBadRequest)
				return
			}
			defer newFile.Close()
			if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
				displayError(w, fmt.Sprintf("Can't write file type: %s", err), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
		} else {
			displayError(w, "Not found", http.StatusNotFound)
		}
	}
}

func handler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error: %v\n%s", r, debug.Stack())
				displayError(w, string(debug.Stack()), http.StatusInternalServerError)
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

func defaultHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		displayError(w, "Not found", http.StatusNotFound)
	}
}

func displayError(w http.ResponseWriter, message string, statusCode int) {
	resp, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

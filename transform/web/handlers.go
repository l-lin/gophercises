package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/l-lin/gophercises/transform/primitive"
)

const (
	maxUploadSize = 100 * 1024 * 1024 // 100Mb
	uploadPath    = "/tmp"
)

// UploadHandler is the handler to upload images
func UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			displayError(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, requestHeader, err := r.FormFile("file")
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

		fileExtensions, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			displayError(w, fmt.Sprintf("Can't read file type: %s", err), http.StatusBadRequest)
			return
		}

		in, err := ioutil.TempFile("", fmt.Sprintf("in_*%s", fileExtensions[0]))
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(in.Name())

		if _, err := in.Write(fileBytes); err != nil {
			displayError(w, fmt.Sprintf("Can't write file type: %s", err), http.StatusBadRequest)
			return
		}

		q := r.URL.Query()
		mode, err := strconv.Atoi(q.Get("mode"))
		if err != nil {
			displayError(w, fmt.Sprintf("Invalid mode: %s", err), http.StatusBadRequest)
			return
		}
		nbShapes, err := strconv.Atoi(q.Get("nbShapes"))
		if err != nil {
			displayError(w, fmt.Sprintf("Invalid nbShapes: %s", err), http.StatusBadRequest)
			return
		}

		out, err := primitive.Transform(in.Name(), fileExtensions[0], mode, nbShapes)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=primitive_%s", requestHeader.Filename))
		w.Header().Add("Content-Type", detectedFileType)
		w.WriteHeader(http.StatusOK)
		if _, err = io.Copy(w, out); err != nil {
			log.Fatal(err)
		}
	}
}

// DefaultHandler to handle errors
func DefaultHandler(h http.Handler) http.HandlerFunc {
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

func displayError(w http.ResponseWriter, message string, statusCode int) {
	log.Printf("[ERROR] %s", message)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": message}); err != nil {
		log.Fatal(err)
	}
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Println(err)
	return true
}

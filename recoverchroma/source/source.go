package source

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// File represents the source code of our project
type File struct {
	Path string
}

// GetFile fetches the source code
func GetFile(path string) (*File, error) {
	if !exists(path) {
		return nil, fmt.Errorf("Could not find the file in path '%s'", path)
	}
	return &File{path}, nil
}

// GetFileName returns the base file name of the file
func (f *File) GetFileName() string {
	return filepath.Base(f.Path)
}

// CopyTo copies the content of the file to the given writer
func (f *File) CopyTo(w io.Writer) error {
	r, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	io.Copy(w, r)
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Println(err)
	return true
}

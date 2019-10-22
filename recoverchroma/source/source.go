package source

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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

// RenderTo copies and highlight code to the given writer
func (f *File) RenderTo(w io.Writer, lineNb int) error {
	r, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	lexer := lexers.Match(f.Path)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}
	ran := [][2]int{
		[2]int{lineNb, lineNb},
	}
	formatter := html.New(html.Standalone(), html.WithLineNumbers(), html.HighlightLines(ran))
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	iterator, err := lexer.Tokenise(nil, string(contents))
	err = formatter.Format(w, style, iterator)
	return err
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

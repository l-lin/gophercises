package renamer

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// NnnRenamer renames files in _NNN
type NnnRenamer struct {
	NbNumbers int
}

// NewNnnRenamer returns a renamer with the pattern _NNN
func NewNnnRenamer(nbNumbers int) *NnnRenamer {
	return &NnnRenamer{nbNumbers}
}

// Rename the given fileName to something like "fileName_NNN.txt"
func (r *NnnRenamer) Rename(nb int, fileName string) string {
	ext := filepath.Ext(fileName)
	fileNameWithoutExt := string([]rune(fileName)[0 : len(fileName)-len(ext)])
	transformedNb, err := r.transform(nb)
	if err != nil {
		log.Fatal(err)
	}
	if fileNameWithoutExt == "" {
		return fmt.Sprintf("%s%s", transformedNb, ext)
	}
	return fmt.Sprintf("%s_%s%s", fileNameWithoutExt, transformedNb, ext)
}

// transform the given number to something like 00XX
// returns an error if the given number is > nb numbers
func (r *NnnRenamer) transform(nb int) (string, error) {
	nbStr := strconv.Itoa(nb)
	if len(nbStr) > r.NbNumbers {
		return "", fmt.Errorf("the given number %d is bigger than the number of numbers %d", nb, r.NbNumbers)
	}
	var b strings.Builder
	for i := 0; i < r.NbNumbers-len(nbStr); i++ {
		b.WriteString("0")
	}
	b.WriteString(nbStr)
	return b.String(), nil
}

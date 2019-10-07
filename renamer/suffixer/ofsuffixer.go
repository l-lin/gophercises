package suffixer

import (
	"log"
	"regexp"
	"strconv"
)

// OfSuffixer extracts fileName suffix (XX of XX)
type OfSuffixer struct {
	Max int
}

// NewOfSuffixer returns the suffixer with the pattern (X of X)
func NewOfSuffixer(max int) *OfSuffixer {
	return &OfSuffixer{max}
}

// Extract the base, nb and ext from a given fileName
func (s *OfSuffixer) Extract(fileName string) (base, ext string, nb int) {
	r := regexp.MustCompile(`(.*)? \((\d+) of \d+\)(\..+)?`)
	if r.Match([]byte(fileName)) {
		result := r.FindAllStringSubmatch(fileName, -1)
		base = result[0][1]
		var err error
		if nb, err = strconv.Atoi(result[0][2]); err != nil {
			log.Fatal(err)
		}
		ext = result[0][3]
		return
	}
	r = regexp.MustCompile(`(.*)(\..+)`)
	if r.Match([]byte(fileName)) {
		result := r.FindAllStringSubmatch(fileName, -1)
		base = result[0][1]
		ext = result[0][2]
		return
	}
	base = fileName
	return
}

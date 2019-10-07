package suffixer

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// NnnSuffixer extracts fileName suffix _NNN
type NnnSuffixer struct {
	NbNumbers int
}

// NewNnnSuffixer returns the suffixer with the pattern _NNN
func NewNnnSuffixer(nbNumbers int) *NnnSuffixer {
	return &NnnSuffixer{nbNumbers}
}

// Extract the base, nb and ext from a given fileName
func (s *NnnSuffixer) Extract(fileName string) (base, ext string, nb int) {
	r := regexp.MustCompile(fmt.Sprintf(`(.*)?_([0-9]{%d})(\..+)?$`, s.NbNumbers))
	if r.Match([]byte(fileName)) {
		result := r.FindAllStringSubmatch(fileName, -1)
		base = result[0][1]
		var err error
		if nb, err = strconv.Atoi(result[0][2]); err != nil {
			log.Fatalln(err)
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

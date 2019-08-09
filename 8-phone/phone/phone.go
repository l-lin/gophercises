package phone

import (
	"regexp"
)

var phoneRexep = regexp.MustCompile("\\D")

//var phoneRexep = regexp.MustCompile(`[0-9]`)

// Phone not normalized fetched from DB
type Phone struct {
	ID    int
	Value string
}

// Normalize phone to expected pattern
func (p *Phone) Normalize() string {
	return phoneRexep.ReplaceAllString(p.Value, "")
}

//func (p *Phone) Normalize() string {
//	var b strings.Builder
//	for _, c := range p.Value {
//		if phoneRexep.MatchString(string(c)) {
//			b.WriteString(string(c))
//		}
//	}
//	return b.String()
//}

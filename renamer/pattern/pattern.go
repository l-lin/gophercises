package pattern

import "strings"

// Getter fetches the pattern associated
type Getter interface {
	GetPatternName() string
}

// String transforms the given getters to a comma separated of patterns
func String(getters ...Getter) string {
	patterns := []string{}
	for _, g := range getters {
		patterns = append(patterns, g.GetPatternName())
	}
	return strings.Join(patterns, ",")
}

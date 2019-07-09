package mapper

import (
	"encoding/json"
)

// FromJSON builds the map from a JSON file
func FromJSON(jsonFile string) (map[string]string, error) {
	return fromFile(jsonFile, json.Unmarshal)
}

package story

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	s, err := ReadFromFile("../cyoa.json")
	if err != nil {
		t.Error(err)
	}
	if s == nil {
		t.Error("Could not parse file.")
	}
	intro := s["intro"]
	if intro == nil {
		t.Error("Intro should not be nil")
	}
}

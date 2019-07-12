package story

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	story, err := ReadFromFile("../cyoa.json")
	if err != nil {
		t.Error(err)
	}
	if story == nil {
		t.Error("Could not parse file.")
	}
	intro := story["intro"]
	if intro == nil {
		t.Error("Intro should not be nil")
	}
}

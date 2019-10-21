package source

import (
	"bytes"
	"testing"
)

const path = "/home/llin/perso/gophercises/recoverchroma/source/source.go"

func TestCopyTo(t *testing.T) {
	f, err := GetFile(path)
	if err != nil {
		t.Errorf("Could not find file '%s'. Error was: %s", path, err)
	}
	w := bytes.NewBufferString("")
	err = f.CopyTo(w)
	if err != nil {
		t.Errorf("Could not copy content of file '%s'. Error was: %s", path, err)
	}
	if len(w.String()) == 0 {
		t.Errorf("Nothing was copied")
	}
}

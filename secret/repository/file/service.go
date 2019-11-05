package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/l-lin/gophercises/secret/secret"
)

// Repository to set the secret in a file
type Repository struct {
	FilePath string
}

// Set the secret in a file
func (r *Repository) Set(s *secret.Secret) error {
	var m map[string]string
	if exists(r.FilePath) {
		b, err := ioutil.ReadFile(r.FilePath)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(b, &m); err != nil {
			return err
		}
	} else {
		m = make(map[string]string, 0)
	}
	m[s.Key] = s.CipherHex
	result, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ioutil.WriteFile(r.FilePath, result, 0600)
	return nil
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Println(err)
	return true
}

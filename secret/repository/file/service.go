package file

import (
	"encoding/json"
	"fmt"
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

// Get the secret from the given key and encoding key
func (r *Repository) Get(key string) (*secret.Secret, error) {
	if !exists(r.FilePath) {
		return nil, fmt.Errorf("%s not found", r.FilePath)
	}

	b, err := ioutil.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}
	var m map[string]string
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	cipherHex, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("No secret found for key %s", key)
	}
	return &secret.Secret{
		Key:       key,
		CipherHex: cipherHex,
	}, nil
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

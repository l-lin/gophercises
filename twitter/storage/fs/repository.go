package fs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/l-lin/gophercises/twitter/user"
)

// Repository to manage users with files
type Repository struct {
	FilePath string
}

// FindAll users from a file
func (r *Repository) FindAll() []user.User {
	dat, err := ioutil.ReadFile(r.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	var users []user.User
	if err = json.Unmarshal(dat, &users); err != nil {
		log.Fatal(err)
	}
	return users
}

// SaveAll users to a file
func (r *Repository) SaveAll(users []user.User) {
	d, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(r.FilePath, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

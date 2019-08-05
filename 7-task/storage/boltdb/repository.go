package boltdb

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/l-lin/7-task/task"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var tasksBucket = []byte("tasks")

// getPath from config file
func getPath() string {
	path := viper.GetString("boltdb.path")
	if path == "" {
		return "task.boltdb"
	}
	return path
}

// Storage in boltdb
type Storage struct {
	path string
}

// NewStorage instanciates a new boltdb repository
func NewStorage() *Storage {
	return &Storage{
		path: getPath(),
	}
}

// GetAll from memory
func (s *Storage) GetAll() []*task.Task {
	db, err := bolt.Open(s.path, 0666, nil)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	tasks := []*task.Task{}
	err = db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return fmt.Errorf("Could not create bucket, error was: %v", err)
		}
		if err = b.ForEach(func(k, v []byte) error {
			var t *task.Task
			if err := json.Unmarshal(v, &t); err != nil {
				return fmt.Errorf("Could not unmarshal task, error was: %v", err)
			}
			tasks = append(tasks, t)
			return nil
		}); err != nil {
			return fmt.Errorf("Could not loop through the bucket, error was: %v", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return tasks
}

// GetIncompletes tasks from boltdb
func (s *Storage) GetIncompletes() []*task.Task {
	return nil
}

// GetCompleted tasks from boltdb
func (s *Storage) GetCompleted() []*task.Task {
	return nil
}

// Add a new task
func (s *Storage) Add(t *task.Task) {
	db, err := bolt.Open(s.path, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return fmt.Errorf("Could not create bucket, error was: %v", err)
		}
		id, _ := b.NextSequence()
		t.ID = int(id)
		buf, err := json.Marshal(t)
		if err != nil {
			return fmt.Errorf("Could not marshal task, error was: %v", err)
		}
		return b.Put(itob(t.ID), buf)
	}); err != nil {
		log.Fatal(err)
	}
}

// Do a task
func (s *Storage) Do(id int) {
	log.WithField("id", id).Warn("No task found for id")
}

// Remove a task from boltdb
func (s *Storage) Remove(id int) {
	log.WithField("id", id).Warn("No task found for id")
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

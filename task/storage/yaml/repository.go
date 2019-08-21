package yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/l-lin/task/internal"
	"github.com/l-lin/task/task"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var tasksBucket = []byte("tasks")

// getPath from config file
func getPath() string {
	path := viper.GetString("yaml.path")
	if path == "" {
		return "task.yaml"
	}
	return path
}

// Storage in yaml file
type Storage struct {
	path string
}

// NewStorage instanciates a new boltdb repository
func NewStorage() *Storage {
	return &Storage{
		path: getPath(),
	}
}

// GetAll tasks from yaml file
func (s *Storage) GetAll() []*task.Task {
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		return nil
	}
	content, err := ioutil.ReadFile(s.path)
	if err != nil {
		log.WithFields(log.Fields{
			"err":  err,
			"file": s.path,
		}).Fatal("Could not read file")
	}
	var out []*task.Task
	err = yaml.Unmarshal(content, &out)
	if err != nil {
		log.WithField("err", err).Fatal("Could not unmarshal file")
	}
	return out
}

// GetIncompletes tasks from yaml file
func (s *Storage) GetIncompletes() []*task.Task {
	return s.filterTask(func(t *task.Task) bool {
		return !t.Completed
	})
}

// GetCompleted tasks from yaml file
func (s *Storage) GetCompleted() []*task.Task {
	return s.filterTask(func(t *task.Task) bool {
		return t.Completed && internal.SameDay(*t.CompletedTime, time.Now())
	})
}

// Add a new task
func (s *Storage) Add(t *task.Task) {
	tasks := s.GetAll()
	t.ID = nextID(tasks)
	tasks = append(tasks, t)
	writeToFile(s.path, tasks)
}

// Do a task
func (s *Storage) Do(id int) {
	tasks := s.GetAll()
	for _, t := range tasks {
		if t.ID == id {
			t.Completed = true
			now := time.Now()
			t.CompletedTime = &now
			fmt.Printf("You have completed the \"%s\" task.\n", t.Content)
			writeToFile(s.path, tasks)
			return
		}
	}
	log.WithField("id", id).Error("Task not found")
}

// Remove a task from the YAML file
func (s *Storage) Remove(id int) {
	tasks := s.GetAll()
	for i, t := range tasks {
		if t.ID == id {
			fmt.Printf("You have deleted the \"%s\" task.\n", t.Content)
			tasks = append(tasks[:i], tasks[i+1:]...)
			writeToFile(s.path, tasks)
			return
		}
	}
	log.WithField("id", id).Error("Task not found")
}

func (s *Storage) filterTask(predicate func(t *task.Task) bool) []*task.Task {
	tasks := s.GetAll()
	filteredTasks := []*task.Task{}
	for _, t := range tasks {
		if predicate(t) {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}

func nextID(tasks []*task.Task) int {
	id := 0
	for _, t := range tasks {
		if id < t.ID {
			id = t.ID
		}
	}
	return id + 1
}

func writeToFile(path string, tasks []*task.Task) {
	content, err := yaml.Marshal(tasks)
	if err != nil {
		log.WithField("err", err).Fatal("Could not marshal tasks")
	}
	err = ioutil.WriteFile(path, content, 0666)
	if err != nil {
		log.WithFields(log.Fields{
			"err":  err,
			"file": path,
		}).Fatal("Could not write file")
	}
}

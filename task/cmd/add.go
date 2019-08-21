package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/l-lin/task/add"
	"github.com/l-lin/task/storage/boltdb"
	"github.com/l-lin/task/storage/yaml"
	"github.com/l-lin/task/task"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run:   runAdd,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) {
	var r add.Repository
	if storageMode == "yaml" {
		r = yaml.NewStorage()
	} else {
		r = boltdb.NewStorage()
	}
	s := add.NewService(r)
	t := &task.Task{
		Content: strings.Join(args, " "),
		Created: time.Now(),
	}
	s.Add(t)
	fmt.Printf("Task \"%s\" added\n", t.Content)
}

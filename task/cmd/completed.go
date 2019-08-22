package cmd

import (
	"fmt"

	"github.com/l-lin/gophercises/task/complete"
	"github.com/l-lin/gophercises/task/storage/boltdb"
	"github.com/l-lin/gophercises/task/storage/yaml"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Show today's completed tasks",
	Run:   runCompleted,
}

func init() {
	rootCmd.AddCommand(completedCmd)
}

func runCompleted(cmd *cobra.Command, args []string) {
	var r complete.Repository
	if storageMode == "yaml" {
		r = yaml.NewStorage()
	} else {
		r = boltdb.NewStorage()
	}
	s := complete.NewService(r)
	tasks := s.GetCompleted()
	if tasks == nil || len(tasks) == 0 {
		fmt.Println("No completed tasks today...")
	} else {
		fmt.Println("You have finished the following tasks today:")
		for _, t := range tasks {
			fmt.Println(t.String())
		}
	}
}

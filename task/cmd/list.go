package cmd

import (
	"fmt"

	"github.com/l-lin/task/list"
	"github.com/l-lin/task/storage/boltdb"
	"github.com/l-lin/task/storage/yaml"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run:   runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) {
	var r list.Repository
	if storageMode == "yaml" {
		r = yaml.NewStorage()
	} else {
		r = boltdb.NewStorage()
	}
	s := list.NewService(r)
	tasks := s.GetIncompletes()
	if tasks == nil || len(tasks) == 0 {
		fmt.Println("Every tasks are complete! You're the best!")
	} else {
		fmt.Println("You have the following tasks:")
		for _, t := range tasks {
			fmt.Println(t.String())
		}
	}
}

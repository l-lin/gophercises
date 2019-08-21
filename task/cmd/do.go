package cmd

import (
	"strconv"

	"github.com/l-lin/task/do"
	"github.com/l-lin/task/storage/boltdb"
	"github.com/l-lin/task/storage/yaml"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run:   runDo,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)
		if err != nil {
			return err
		}
		_, err = strconv.Atoi(args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

func runDo(cmd *cobra.Command, args []string) {
	var r do.Repository
	if storageMode == "yaml" {
		r = yaml.NewStorage()
	} else {
		r = boltdb.NewStorage()
	}
	s := do.NewService(r)
	id, _ := strconv.Atoi(args[0])
	s.Do(id)
}

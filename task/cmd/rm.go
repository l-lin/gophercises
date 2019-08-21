package cmd

import (
	"strconv"

	"github.com/l-lin/task/rm"
	"github.com/l-lin/task/storage/boltdb"
	"github.com/l-lin/task/storage/yaml"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task",
	Run:   runRm,
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
	rootCmd.AddCommand(rmCmd)
}

func runRm(cmd *cobra.Command, args []string) {
	var r rm.Repository
	if storageMode == "yaml" {
		r = yaml.NewStorage()
	} else {
		r = boltdb.NewStorage()
	}
	s := rm.NewService(r)
	id, _ := strconv.Atoi(args[0])
	s.Remove(id)
}

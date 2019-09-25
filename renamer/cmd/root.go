package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "renamer",
		Short: "File renaming tool",
		Long:  "Rename files with the given naming pattern",
		Args:  cobra.ExactArgs(1),
		Run:   run,
	}
	pattern string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Naming pattern")
}

func run(cmd *cobra.Command, args []string) {
}

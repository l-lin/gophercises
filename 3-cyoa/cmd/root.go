package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	inputFile string
	startArc  string
	rootCmd   = &cobra.Command{
		Use:   "cyoa",
		Short: "Choose Your Own Adventure",
		Long: `Choose Your Own Adventure is a series of children's gamebooks where each story is written
from a second-person point of view, with the reader assuming the role of the protagonist and making
choices that determine the main character's actions and the plot's outcome.`,
	}
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
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "cyoa.json", "input JSON story file")
	rootCmd.PersistentFlags().StringVar(&startArc, "start", "intro", "arc starter of the story")
}

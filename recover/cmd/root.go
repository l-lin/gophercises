package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	devMode bool
	rootCmd = &cobra.Command{
		Use:   "recover",
		Short: "Panic & recover mechanisms exercises",
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
	rootCmd.PersistentFlags().BoolVar(&devMode, "dev", false, "development mode")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "secret",
		Short: "CLI that manages secrets",
	}
	filePath    string
	encodingKey string
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
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&filePath, "file-path", "f", fmt.Sprintf("%s/%s.json", path, "secrets"), "file that will contains the secrets")
	rootCmd.PersistentFlags().StringVarP(&encodingKey, "encoding-key", "k", "", "encoding key to encrypt the secret")
	rootCmd.MarkFlagRequired("encoding-key")
}

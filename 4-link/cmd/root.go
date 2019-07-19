package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/l-lin/4-link/link"
	"github.com/spf13/cobra"
)

var inputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "link",
	Short: "HTML Link Parser",
	Long:  "Extracts link by parsing given HTML file",
	Run:   run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	links, err := link.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(links)
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input HTML file to parse")
	rootCmd.MarkFlagRequired("input")
}

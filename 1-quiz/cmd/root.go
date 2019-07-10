package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/l-lin/1-quiz/problem"
	"github.com/l-lin/1-quiz/query"
	"github.com/spf13/cobra"
)

var (
	inputFile string
	timer     time.Duration
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Execute a quiz based on a CSV input file",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "problems.csv", "input CSV file (default is ./problems.csv)")
	rootCmd.PersistentFlags().DurationVarP(&timer, "time-limit", "t", 30*time.Second, "time limit to answer the questions")
}

func run(cmd *cobra.Command, args []string) {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Could not read file %s, error was: %v", inputFile, err)
	}
	pbParser := &problem.CsvProblemsParser{}
	pbs, err := pbParser.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatalf("Could not parse the file %s, error was: %v", inputFile, err)
	}

	querier := &query.ConsoleQuerier{}
	ready, err := querier.AskReady()
	if err != nil {
		log.Fatalln(err)
	}
	if ready {
		result, err := query.Query(querier, pbs, timer)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(result)
	}
}

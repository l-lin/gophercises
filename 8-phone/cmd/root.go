package cmd

import (
	"fmt"
	"os"

	"github.com/l-lin/8-phone/list"
	"github.com/l-lin/8-phone/storage/postgresql"
	"github.com/spf13/cobra"
)

var (
	dbURL   string
	rootCmd = &cobra.Command{
		Use:   "phone",
		Short: "Iterate through a database a normalize all phone numbers",
		Run:   run,
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

func run(cmd *cobra.Command, args []string) {
	r := postgresql.New(dbURL)
	s := list.NewService(r)
	phones := s.GetAll()
	for _, p := range phones {
		fmt.Println(p.Value)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbURL, "db-url", "postgres://postgres@localhost:5432/phone?sslmode=disable", "DB URL")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/l-lin/8-phone/list"
	"github.com/l-lin/8-phone/phone"
	"github.com/l-lin/8-phone/rm"
	"github.com/l-lin/8-phone/storage/postgresql"
	"github.com/l-lin/8-phone/update"
	"github.com/spf13/cobra"
)

var (
	dbURL   string
	rootCmd = &cobra.Command{
		Use:   "phone",
		Short: "Iterate through a database and normalize all phone numbers",
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
	listService := list.NewService(r)
	updateService := update.NewService(r)
	deleteService := rm.NewService(r)

	phones := listService.GetAll()
	uniquePhones := make([]*phone.Phone, 0)
	m := make(map[string]bool, 0)
	for _, p := range phones {
		p.Value = p.Normalize()
		if listService.Count(p.Value) > 1 {
			deleteService.Delete(p.ID)
		} else {
			updateService.Update(p)
		}
		if _, ok := m[p.Value]; !ok {
			uniquePhones = append(uniquePhones, p)
			m[p.Value] = true
		}
	}

	for _, p := range uniquePhones {
		fmt.Println(p.Value)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbURL, "db-url", "postgres://postgres:postgres@localhost:5433/phone?sslmode=disable", "DB URL")
}

package cmd

import (
	"fmt"
	"log"

	"github.com/l-lin/gophercises/secret/encrypt"
	"github.com/l-lin/gophercises/secret/get"
	"github.com/l-lin/gophercises/secret/repository/file"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the secret from the given key",
		Args:  cobra.ExactArgs(1),
		Run:   runGet,
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGet(cmd *cobra.Command, args []string) {
	r := &file.Repository{FilePath: filePath}
	service := get.New(r)
	s, err := service.Get(args[0])
	if err != nil {
		log.Fatal(err)
	}
	result, err := encrypt.Decrypt(encodingKey, s.CipherHex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

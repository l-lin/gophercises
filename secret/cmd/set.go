package cmd

import (
	"log"

	"github.com/l-lin/gophercises/secret/encrypt"
	"github.com/l-lin/gophercises/secret/repository/file"
	"github.com/l-lin/gophercises/secret/secret"
	"github.com/l-lin/gophercises/secret/set"
	"github.com/spf13/cobra"
)

var (
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set the secret",
		Args:  cobra.ExactArgs(2),
		Run:   run,
	}
	encodingKey string
)

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&encodingKey, "encoding-key", "k", "", "encoding key to encrypt the secret")
	setCmd.MarkFlagRequired("encoding-key")
}

func run(cmd *cobra.Command, args []string) {
	r := &file.Repository{FilePath: filePath}
	service := set.New(r)
	cipherHex, err := encrypt.Encrypt(encodingKey, args[1])
	if err != nil {
		log.Fatal(err)
	}

	s := &secret.Secret{
		Key:       args[0],
		CipherHex: cipherHex,
	}

	err = service.Set(s)
	if err != nil {
		log.Fatal(err)
	}
}

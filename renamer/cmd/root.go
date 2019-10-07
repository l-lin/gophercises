package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/l-lin/gophercises/renamer/renamer"
	"github.com/l-lin/gophercises/renamer/suffixer"
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
	from, to  string
	dry       bool
	renamers  = make(map[string]renamer.Renamer)
	suffixers = make(map[string]suffixer.Suffixer)
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
	// TODO: set max & nbNumbers by the user instead
	renamers["of"] = renamer.NewOfRenamer(100)
	renamers["nnn"] = renamer.NewNnnRenamer(3)
	suffixers["of"] = suffixer.NewOfSuffixer(100)
	suffixers["nnn"] = suffixer.NewNnnSuffixer(3)
	rootCmd.Flags().BoolVar(&dry, "dry", false, "dry run")
	rootCmd.Flags().StringVarP(
		&from,
		"from",
		"f",
		"",
		fmt.Sprintf("From naming pattern (available patterns: %v)", getPatterns(reflect.ValueOf(suffixers).MapKeys())),
	)
	rootCmd.Flags().StringVarP(
		&to,
		"to",
		"t",
		"",
		fmt.Sprintf("To naming pattern (available patterns: %v)", getPatterns(reflect.ValueOf(renamers).MapKeys())),
	)
	rootCmd.MarkFlagRequired("from")
	rootCmd.MarkFlagRequired("to")
}

func run(cmd *cobra.Command, args []string) {
	s := suffixers[from]
	if s == nil {
		log.Fatalf("no suffixer found for type '%s'", from)
	}
	r := renamers[to]
	if r == nil {
		log.Fatalf("no renamer found for type '%s'", to)
	}
	path := args[0]
	if !exists(path) {
		log.Fatalf("given path %s does not exist", path)
	}
	rename(path, s, r)
}

func getPatterns(values []reflect.Value) string {
	patterns := make([]string, len(values))
	for i, v := range values {
		patterns[i] = v.String()
	}
	return strings.Join(patterns, ",")
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Println(err)
	return true
}

func rename(path string, s suffixer.Suffixer, r renamer.Renamer) {
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		renameFile(path, s, r)
	case mode.IsDir():
		err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				renameFile(p, s, r)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	case mode&os.ModeSymlink != 0:
		log.Printf("%s is a symlink and it's not handled by this app\n", path)
	}
}

func renameFile(filePath string, s suffixer.Suffixer, r renamer.Renamer) {
	base, ext, nb := s.Extract(filePath)
	if nb == 0 {
		// case where filePath is not supported by the suffixer type
		// maybe return something else to understand it's not supported instead of check nb...
		return
	}
	newFilePath := r.Rename(nb, base+ext)
	if newFilePath == "" {
		log.Printf("could not rename file '%s' with type '%s'\n", filePath, to)
		return
	}
	fmt.Printf("%s => %s\n", filePath, newFilePath)
	if !dry {
		if err := os.Rename(filePath, newFilePath); err != nil {
			log.Fatalln(err)
		}
	}
}

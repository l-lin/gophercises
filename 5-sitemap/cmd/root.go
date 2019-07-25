package cmd

import (
	"os"
	"strconv"

	"github.com/l-lin/5-sitemap/crawler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const maxDepth = 3

var (
	depth   int
	rootCmd = &cobra.Command{
		Use:   "sitemap",
		Short: "Build a sitemap",
		Long:  `A sitemap is basically a map of all of the pages within a specific domain. They are used by search engines and other tools to inform them of all of the pages on your domain.`,
		Run:   run,
		Args:  cobra.ExactArgs(1),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&depth, "depth", "d", maxDepth, "Maximum number of links to follow")
}

func run(cmd *cobra.Command, args []string) {
	c, err := crawler.New(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	links, err := c.Perform(depth)
	if err != nil {
		log.Fatalln(err)
	}
	fields := log.Fields{}
	for i, l := range links {
		fields[strconv.Itoa(i)] = l.String()
	}
	log.WithFields(fields).Info("Links fetched")
}

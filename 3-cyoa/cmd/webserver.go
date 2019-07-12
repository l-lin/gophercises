package cmd

import "github.com/spf13/cobra"

var webserverCmd = &cobra.Command{
	Use:   "webserver",
	Short: "Start the story in a web server",
	Run:   executeWebServer,
}

func init() {
	rootCmd.AddCommand(webserverCmd)
}

func executeWebServer(cmd *cobra.Command, args []string) {
}

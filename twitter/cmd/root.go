package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/l-lin/commuting-traffic-info/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const cfgFileName = ".twitter-contest"

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "twitter",
		Short: "A twitter contest CLI to determine who is the winner from retweeters",
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
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s.yaml)", cfgFileName))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".commuting-traffic-info" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgFileName)
		cfgFile = fmt.Sprintf("%s/%s.yaml", home, cfgFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		key := config.GetAPIKey()
		secretKey := config.GetAPISecretKey()
		if key == "" || secretKey == "" {
			log.Println("Could not read the 'key' and 'secret-key' properties. Initializing it")
			config.InitTwitterAPIKeys(cfgFile)
		}
	} else {
		config.InitTwitterAPIKeys(cfgFile)
	}
}

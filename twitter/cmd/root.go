package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/l-lin/gophercises/twitter/config"
	"github.com/l-lin/gophercises/twitter/storage/fs"
	"github.com/l-lin/gophercises/twitter/twitter"
	"github.com/l-lin/gophercises/twitter/user"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const cfgFileName = ".twitter-contest"

var (
	cfgFile    string
	pickWinner bool
	rootCmd    = &cobra.Command{
		Use:   "twitter",
		Short: "A twitter contest CLI to determine who is the winner from retweeters",
		Run:   run,
		Args:  cobra.MinimumNArgs(1),
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
	tweetID := args[0]
	r := &fs.Repository{FilePath: "./users.json"}
	s := user.NewService(r)
	users := s.FindAll()
	retweetUsers := getRetweetUsers(tweetID)
	users = user.Merge(users, retweetUsers)
	s.SaveAll(users)
	if pickWinner {
		winner := s.PickWinner(users)
		fmt.Println("And the winner is...", winner.Name)
	}
}

func getRetweetUsers(tweetID string) []user.User {
	retweetsResultCh := make(chan *twitter.RetweetsResult)
	go twitter.GetRetweets(retweetsResultCh, tweetID)
	result := <-retweetsResultCh
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	return result.GetUniqueUsers()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s.yaml)", cfgFileName))
	rootCmd.PersistentFlags().BoolVar(&pickWinner, "pick-winner", false, "pick winner among the retweeters")
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

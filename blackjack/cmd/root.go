package cmd

import (
	"fmt"
	"os"

	"github.com/l-lin/gophercises/blackjack/game"
	"github.com/spf13/cobra"
)

var nbPlayers int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Simple blackjack game",
	Run:   run,
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
	rootCmd.Flags().IntVarP(&nbPlayers, "players", "p", 1, "number of players")
}

func run(cmd *cobra.Command, args []string) {
	//	fmt.Print(`
	//  , _
	// /|/_) |\   _,    _   |)    o   _,    _   |)
	//  |  \ |/  / |   /    |/)   |  / |   /    |/)
	//  |(_/ |_/ \/|_/ \__/ | \/  |/ \/|_/ \__/ | \/
	//                           (|
	//`)
	g := game.New(nbPlayers)
	g.Run()
}

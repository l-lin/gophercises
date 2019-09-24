package cmd

import (
	"fmt"
	"os"

	"github.com/l-lin/gophercises/blackjack/game"
	"github.com/spf13/cobra"
)

const maxNbPlayers = 2

var (
	nbPlayers int
	rootCmd   = &cobra.Command{
		Use:   "blackjack",
		Short: "Simple blackjack game",
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
	if nbPlayers > maxNbPlayers {
		fmt.Printf("There can only have %d players\n", maxNbPlayers)
		os.Exit(1)
	}
	g := game.New(nbPlayers)
	g.Run()
}

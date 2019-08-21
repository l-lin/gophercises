package cmd

import (
	"fmt"
	"log"

	"github.com/l-lin/cyoa/story"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Start the story in the console",
	Run:   executeConsole,
}

func init() {
	rootCmd.AddCommand(consoleCmd)
}

func executeConsole(cmd *cobra.Command, args []string) {
	s, err := story.ReadFromFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	display(s, startArc)
}

func display(s map[string]*story.Chapter, arc string) {
	chapter, ok := s[arc]
	if !ok {
		log.Fatalf("Could not find the arc %s from the story", arc)
	}

	fmt.Printf("\n%s\n\n", chapter.Title)
	for _, content := range chapter.Story {
		fmt.Println(content)
	}

	items := make([]string, 0)
	for _, option := range chapter.Options {
		items = append(items, option.Text)
	}

	if chapter.Options == nil || len(chapter.Options) == 0 {
		return
	}
	prompt := promptui.Select{
		Label: "Choose your next step",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalln(err)
	}

	var arcChoice string
	for _, option := range chapter.Options {
		if result == option.Text {
			arcChoice = option.Arc
			break
		}
	}
	display(s, arcChoice)
}

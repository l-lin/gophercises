package cmd

import (
	"github.com/l-lin/gophercises/image/charts"
	"github.com/spf13/cobra"
)

var (
	chartsCmd = &cobra.Command{
		Use:   "charts",
		Short: "draw charts using golang standard image library",
		Run:   runCharts,
	}
	filename string
)

func init() {
	rootCmd.AddCommand(chartsCmd)
	chartsCmd.Flags().StringVarP(&filename, "output", "o", "", "image file name to output the result")
	chartsCmd.MarkFlagRequired("output")
}

func runCharts(cmd *cobra.Command, args []string) {
	data := []int{10, 33, 73, 64}
	charts.Draw(filename, data)
}

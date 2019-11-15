package cmd

import (
	"log"

	"github.com/l-lin/gophercises/image/charts"
	"github.com/l-lin/gophercises/image/mask"
	"github.com/l-lin/gophercises/image/pixelbypixel"
	"github.com/spf13/cobra"
)

var (
	chartsCmd = &cobra.Command{
		Use:   "charts",
		Short: "draw charts using golang standard image library",
		Run:   runCharts,
	}
	filename   string
	drawerType string
)

func init() {
	rootCmd.AddCommand(chartsCmd)
	chartsCmd.Flags().StringVarP(&filename, "output", "o", "", "image file name to output the result")
	chartsCmd.Flags().StringVarP(&drawerType, "drawer", "d", "pixelbypixel", "drawer to use to draw the charts (one of: pixelbypixel, mask)")
	chartsCmd.MarkFlagRequired("output")
}

func runCharts(cmd *cobra.Command, args []string) {
	data := []int{10, 33, 73, 64}
	var drawer charts.Drawer
	if "pixelbypixel" == drawerType {
		drawer = &pixelbypixel.Drawer{}
	} else if "mask" == drawerType {
		drawer = &mask.Drawer{}
	} else {
		log.Fatalf("no drawer type found for '%s'", drawerType)
	}
	charts.Draw(drawer, filename, data)
}

package charts

import (
	"image"
	"image/png"
	"os"
)

const (
	// BarWidth is the width of the chart bars
	BarWidth = 50
	// SepWidth is the width of the seperator between the chart bars
	SepWidth = 10
	// BarHeightCoeff is the coefficient to use to multiply the chart bar heights
	BarHeightCoeff = 2
)

// Drawer draws the charts using different techniques
type Drawer interface {
	Draw(w, h int, data []int) (image.Image, error)
}

// Draw charts using standard image package
func Draw(drawer Drawer, filename string, data []int) error {
	w := computeMaxWidth(data)
	h := computeMaxHeight(data)
	img, err := drawer.Draw(w, h, data)
	if err != nil {
		return err
	}
	return writeToFile(filename, img)
}

func computeMaxHeight(data []int) int {
	max := 0
	for _, d := range data {
		if max < d {
			max = d
		}
	}
	return max * BarHeightCoeff
}

func computeMaxWidth(data []int) int {
	if len(data) == 0 {
		return 0
	}
	nbSep := len(data) - 1
	if len(data) < 1 {
		nbSep = 0
	}
	return BarWidth*len(data) + SepWidth*nbSep
}

func writeToFile(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)
	return nil
}

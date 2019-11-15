package pixelbypixel

import (
	"image"
	"image/color"

	"github.com/l-lin/gophercises/image/charts"
)

// Drawer draws the charts pixel by pixel
type Drawer struct {
}

// Draw the charts pixel by pixel
func (d *Drawer) Draw(w, h int, data []int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.SetRGBA(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	for i, d := range data {
		startX := i*charts.BarWidth + i*charts.SepWidth
		endX := startX + charts.BarWidth
		for x := startX; x < endX; x++ {
			for y := h; y >= (h - d*charts.BarHeightCoeff); y-- {
				img.SetRGBA(x, y, color.RGBA{0, 0, 255, 255})
			}
		}
	}
	return img
}

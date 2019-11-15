package mask

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/l-lin/gophercises/image/charts"
)

// Drawer draws the charts using image.draw package with masks
type Drawer struct {
}

// Draw the charts using masks
func (d *Drawer) Draw(w, h int, data []int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bg := image.NewUniform(color.RGBA{255, 255, 255, 255})
	draw.Draw(img, img.Bounds(), bg, image.ZP, draw.Src)

	blue := image.NewUniform(color.RGBA{0, 0, 255, 255})
	for i, d := range data {
		startX := i*charts.BarWidth + i*charts.SepWidth
		endX := startX + charts.BarWidth
		startY := h - d*charts.BarHeightCoeff
		endY := h
		bar := image.Rect(startX, startY, endX, endY)
		draw.Draw(img, bar, blue, image.ZP, draw.Src)
	}

	return img
}

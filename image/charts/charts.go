package charts

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	barWidth       = 50
	sepWidth       = 10
	barHeightCoeff = 2
)

// Draw charts using standard image package
func Draw(filename string, data []int) error {
	w := computeMaxWidth(data)
	h := computeMaxHeight(data)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.SetRGBA(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	for i, d := range data {
		start := i*barWidth + i*sepWidth
		end := start + barWidth
		for x := start; x < end; x++ {
			for y := h; y >= (h - d*barHeightCoeff); y-- {
				img.SetRGBA(x, y, color.RGBA{0, 0, 255, 255})
			}
		}
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
	return max * barHeightCoeff
}

func computeMaxWidth(data []int) int {
	if len(data) == 0 {
		return 0
	}
	nbSep := len(data) - 1
	if len(data) < 1 {
		nbSep = 0
	}
	return barWidth*len(data) + sepWidth*nbSep
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

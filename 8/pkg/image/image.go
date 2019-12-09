package image

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	Layers []Layer
	Width  int
	Height int
	Img    *image.RGBA
}

type Layer struct {
	Data       [][]int
	ZeroDigits int
	OneDigits  int
	TwoDigits  int
}

func New(data []int, width, height int) *Image {
	image := Image{
		Layers: make([]Layer, 0),
		Width:  width,
		Height: height,
	}

	w, h, zerodigits, onedigits, twodigits := 0, 0, 0, 0, 0
	currlayer := make([][]int, 0)
	currwidth := make([]int, 0)
	for _, digit := range data {
		switch digit {
		case 0:
			zerodigits++
		case 1:
			onedigits++
		case 2:
			twodigits++
		}
		currwidth = append(currwidth, digit)
		w++
		if w == width {
			currlayer = append(currlayer, currwidth)
			currwidth = make([]int, 0)
			w = 0
			h++
		}
		if h == height {
			l := Layer{
				Data:       currlayer,
				ZeroDigits: zerodigits,
				OneDigits:  onedigits,
				TwoDigits:  twodigits,
			}
			image.Layers = append(image.Layers, l)
			zerodigits, onedigits, twodigits, w, h = 0, 0, 0, 0, 0
			currlayer = make([][]int, 0)
		}
	}
	return &image
}

func (i *Image) Decode() {
	newImage := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))

	var pixel color.Color
	for x := 0; x < i.Width; x++ {
		for y := 0; y < i.Height; y++ {
			l := 0
			var c int
			pixel = color.Transparent
			for l < len(i.Layers) {
				layer := i.Layers[l]
				c = layer.Data[y][x]
				if c == 1 {
					pixel = color.White
					break
				} else if c == 0 {
					pixel = color.Black
					break
				} else {
					pixel = color.Transparent
				}
				l++
			}
			newImage.Set(x, y, pixel)
		}
	}

	fmt.Println(newImage)
	i.Img = newImage
}

func (i *Image) Write(filename string) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = png.Encode(outputFile, i.Img)
	if err != nil {
		return err
	}
	outputFile.Close()
	return nil
}

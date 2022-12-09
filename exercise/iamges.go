package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{interpretCoordinate(x,y),interpretCoordinate(x,y), 255, 255}
}

func interpretCoordinate(x, y int) uint8 {
	n := x^y
	
	return uint8(n)
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}

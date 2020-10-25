package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Circle struct {
	p image.Point
	r int
}

type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

var blue color.Color = color.RGBA{17, 62, 105, 255}
var red color.Color = color.RGBA{250, 5, 5, 255}

func main() {
	file, err := os.Create("myImage.png")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	mask := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(mask, mask.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{20, 21}, 20}, image.ZP, draw.Over)
	draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{379, 141}, 40}, image.ZP, draw.Over)
	draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{239, 321}, 70}, image.ZP, draw.Over)

	//This code drawing single line
	// draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// for y := 20; y < 380; y++ {
	// 	x := 20
	// 	img.Set(x, y, red)
	// }

	png.Encode(file, img)
}

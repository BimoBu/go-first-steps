/*

Exercise: Images
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

*/

package image

import (
	"fmt"
	"image"
	"image/color"
	// "golang.org/x/tour/pic"
)

type Image struct {
	SizeX, SizeY int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{i.SizeX, i.SizeY}}
}

func (i Image) At(x, y int) color.Color {
	max := 255
	r := uint8((x/2 - y) % max)
	g := uint8((x - y) % max)
	b := uint8((x - y) % max)
	const a uint8 = 128
	return color.RGBA{r, g, b, a}
}

func Test() {
	fmt.Println("Image needs to be executed in the Go Playground")
	// m := Image{255, 255}
	// pic.ShowImage(m)
}

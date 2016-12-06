package halftone

import (
	"image"
	"image/color"
	"math"
)

type ImageConverter struct {
}

func NewImageConverter() *ImageConverter {
	return new(ImageConverter)
}

func (ic *ImageConverter) ConvertToGray(m image.Image) *image.Gray {
	return ic.convert(m, ic.toGrayLuminance)
}

type ConvertFunc func(color.Color) color.Gray

// Convert converts a color image to grayscale using the provided conversion function.
func (ic *ImageConverter) convert(m image.Image, convertColor ConvertFunc) *image.Gray {
	b := m.Bounds()
	gray := image.NewGray(b)
	pos := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			gray.Pix[pos] = convertColor(m.At(x, y)).Y
			pos++
		}
	}
	return gray
}

// ToGrayLuminance converts color.Color c to grayscale using Rec 709.
//
// The formula used for conversion is: Y' = 0.2125*R' + 0.7154*G' + 0.0721*B'
// where r, g and b are gamma expanded with gamma 2.2 and final Y is Y'
// gamma compressed again.
// The same formula is used by color.GrayModel.Convert().
func (ic *ImageConverter) toGrayLuminance(c color.Color) color.Gray {
	rr, gg, bb, _ := c.RGBA()
	r := math.Pow(float64(rr), 2.2)
	g := math.Pow(float64(gg), 2.2)
	b := math.Pow(float64(bb), 2.2)
	y := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
	Y := uint16(y + 0.5)
	return color.Gray{uint8(Y >> 8)}
}

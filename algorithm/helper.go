package algorithm

import (
	"github.com/pichuio/halftone/common"
	"image"
)

func ConvertColoredArrToGrayArr(a *common.Array) *common.Array {
	grayArr := common.CloneArray(a)
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {

		}
	}

	return grayArr
}

func ConvertImageToGrayArr(image image.Image) *common.Array {
	bounds := image.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	grayArr := common.NewArray(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			red, green, blue, _ := image.At(x, y).RGBA()
			red2 := uint8(red)
			green2 := uint8(green)
			blue2 := uint8(blue)
			gray := CalculateGray(red2, green2, blue2)
			gray_ := uint8(gray)
			grayArr.Array[x][y] = int32(gray_)
		}
	}

	return grayArr
}

func CalculateGray(red, green, blue uint8) uint32 {
	return (uint32(red) + uint32(green) + uint32(blue)) / 3
}

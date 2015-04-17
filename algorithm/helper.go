package algorithm

import (
	"github.com/pichuio/halftone/common"
	"image"
	"image/color"
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

func ConvertGrayArrayToImage(arr *common.Array) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, arr.Width, arr.Height))
	for y := 0; y < arr.Height; y++ {
		for x := 0; x < arr.Width; x++ {
			//pixel := grayImage.At(y, x)
			temp := arr.Array[x][y]

			c := color.RGBA{uint8(temp), uint8(temp), uint8(temp), 0xff}
			dst.SetRGBA(x, y, c)
		}
	}
	return dst
}

func GetClosestColor(color int32) int32 {
	if color < 128 {
		return 0
	}
	return 255
}

func CalculateGray(red, green, blue uint8) uint32 {
	return (uint32(red) + uint32(green) + uint32(blue)) / 3
}

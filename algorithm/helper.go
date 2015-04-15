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

func DitheringMatrix2x3_1(a *common.Array) *common.Array {
	arr := common.CloneArray(a)
	for x := 0; x < arr.Width; x++ {
		for y := 0; y < arr.Height; y++ {
			oldPixel := arr.Array[x][y]

			if oldPixel < 128 {
				arr.Array[x][y] = 0
			} else {
				arr.Array[x][y] = 255
			}

			quantError := oldPixel - arr.Array[x][y]

			if y > 0 && x+1 < arr.Width {
				arr.Array[x+1][y-1] += int32(float64(quantError) * (3.0 / 16.0))
			}
			if x+1 < arr.Width {
				arr.Array[x+1][y] += int32(float64(quantError) * (5.0 / 16.0))
			}
			if y+1 < arr.Height && x+1 < arr.Width {
				arr.Array[x+1][y+1] += int32(float64(quantError) * (1.0 / 16.0))
			}
			if y+1 < arr.Height {
				arr.Array[x][y+1] += int32(float64(quantError) * (7.0 / 16.0))
			}
		}
	}

	return arr

}

func DitheringMatrix2x3_2(a *common.Array) *common.Array {
	arr := common.CloneArray(a)
	for y := arr.Height - 1; y >= 0; y-- {
		for x := 0; x < arr.Width; x++ {
			oldPixel := arr.Array[x][y]

			if oldPixel < 128 {
				arr.Array[x][y] = 0
			} else {
				arr.Array[x][y] = 255
			}

			quantError := oldPixel - arr.Array[x][y]

			if x > 0 && y > 0 {
				arr.Array[x-1][y-1] += int32(float64(quantError) * (3.0 / 16.0))
			}
			if y > 0 {
				arr.Array[x][y-1] += int32(float64(quantError) * (5.0 / 16.0))
			}
			if y > 0 && x+1 < arr.Width {
				arr.Array[x+1][y-1] += int32(float64(quantError) * (1.0 / 16.0))
			}
			if x+1 < arr.Width {
				arr.Array[x+1][y] += int32(float64(quantError) * (7.0 / 16.0))
			}
		}
	}

	return arr

}

func CalculateGray(red, green, blue uint8) uint32 {
	return (uint32(red) + uint32(green) + uint32(blue)) / 3
}

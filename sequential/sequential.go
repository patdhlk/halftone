package sequential

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image"
)

// Start position: Left Top
// Dither order: Left to right; continue with next row
func Dithering(a *common.Array, mode int, factorErr float64) *common.Array {
	arr := common.CloneArray(a)
	for y := arr.Height - 1; y >= 0; y-- {
		for x := 0; x < arr.Width; x++ {
			oldPixel := arr.Array[x][y]

			arr.Array[x][y] = algorithm.GetClosestColor(oldPixel)

			quantError := oldPixel - arr.Array[x][y]
			factorQuantError := float64(quantError) * factorErr

			//  X  P  7
			//  3  5  1
			if mode == 0 {
				if x+1 < arr.Width {
					arr.Array[x+1][y] += int32(factorQuantError * (7.0 / 16.0))
				}
				if y > 0 {
					if x > 0 {
						arr.Array[x-1][y-1] += int32(factorQuantError * (3.0 / 16.0))
					}
					arr.Array[x][y-1] += int32(factorQuantError * (5.0 / 16.0))
					if x+1 < arr.Width {
						arr.Array[x+1][y-1] += int32(factorQuantError * (1.0 / 16.0))
					}
				}
			}
			//  X  P  4  1
			//  1  4  1
			//     1
			if mode == 1 {
				if x+1 < arr.Width {
					arr.Array[x+1][y] += int32(factorQuantError * (4.0 / 12.0))
				}
				if x+2 < arr.Width {
					arr.Array[x+2][y] += int32(factorQuantError * (1.0 / 12.0))
				}
				if y > 0 && x > 0 {
					arr.Array[x-1][y-1] += int32(factorQuantError * (1.0 / 12.0))
				}
				if y > 0 {
					arr.Array[x][y-1] += int32(factorQuantError * (4.0 / 12.0))
				}
				if y > 0 && x+1 < arr.Width {
					arr.Array[x+1][y-1] += int32(factorQuantError * (1.0 / 12.0))
				}
				if y > 1 {
					arr.Array[x][y-2] += int32(factorQuantError * (1.0 / 12.0))
				}
			}

			//  X  X  P  8  4
			//  2  4  8  4  2
			//  1  2  4  2  1
			if mode == 2 {
				if x+1 < arr.Width {
					arr.Array[x+1][y] += int32(factorQuantError * (8.0 / 42.0))
				}
				if x+2 < arr.Width {
					arr.Array[x+2][y] += int32(factorQuantError * (4.0 / 42.0))
				}
				if y > 0 {
					if x > 1 {
						arr.Array[x-2][y-1] += int32(factorQuantError * (2.0 / 42.0))
					}
					if x > 0 {
						arr.Array[x-1][y-1] += int32(factorQuantError * (4.0 / 42.0))
					}
					arr.Array[x][y-1] += int32(factorQuantError * (8.0 / 42.0))
					if x+1 < arr.Width {
						arr.Array[x+1][y-1] += int32(factorQuantError * (4.0 / 42.0))
					}
					if x+2 < arr.Width {
						arr.Array[x+2][y-1] += int32(factorQuantError * (2.0 / 42.0))
					}
				}
				if y > 1 {
					if x > 1 {
						arr.Array[x-2][y-2] += int32(factorQuantError * (1.0 / 42.0))
					}
					if x > 0 {
						arr.Array[x-1][y-2] += int32(factorQuantError * (2.0 / 42.0))
					}
					arr.Array[x][y-2] += int32(factorQuantError * (4.0 / 42.0))
					if x+1 < arr.Width {
						arr.Array[x+1][y-2] += int32(factorQuantError * (2.0 / 42.0))
					}
					if x+2 < arr.Width {
						arr.Array[x+2][y-2] += int32(factorQuantError * (1.0 / 42.0))
					}
				}
			}
		}
	}
	return arr
}

func RunSequentialMain(arr *common.Array, mode int, factorErr float64) *image.RGBA {
	arr = Dithering(arr, mode, factorErr)

	dst := algorithm.ConvertGrayArrayToImage(arr)
	return dst
}

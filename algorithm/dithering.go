package algorithm

import (
	"github.com/pichuio/halftone/common"
)

// Start position: Left button
// Dither order: Bottom to top; continue with next column
// Matrix:
//  7  1
//  P  5
//  X  3
func DitheringMatrix2x3_1(a *common.Array, factorErr float64) *common.Array {
	arr := common.CloneArray(a)
	for x := 0; x < arr.Width; x++ {
		for y := 0; y < arr.Height; y++ {
			oldPixel := arr.Array[x][y]

			arr.Array[x][y] = GetClosestColor(oldPixel)

			quantError := oldPixel - arr.Array[x][y]
			factorQuantError := float64(quantError) * factorErr

			if y > 0 && x+1 < arr.Width {
				arr.Array[x+1][y-1] += int32(factorQuantError * (3.0 / 16.0))
			}
			if x+1 < arr.Width {
				arr.Array[x+1][y] += int32(factorQuantError * (5.0 / 16.0))
			}
			if y+1 < arr.Height && x+1 < arr.Width {
				arr.Array[x+1][y+1] += int32(factorQuantError * (1.0 / 16.0))
			}
			if y+1 < arr.Height {
				arr.Array[x][y+1] += int32(factorQuantError * (7.0 / 16.0))
			}
		}
	}
	return arr
}

// Start position: Left Top
// Dither order: Left to right; continue with next row
// Matrix:
//  X  P  7
//  3  5  1
func DitheringMatrix2x3_2(a *common.Array, factorErr float64) *common.Array {
	arr := common.CloneArray(a)
	for y := arr.Height - 1; y >= 0; y-- {
		for x := 0; x < arr.Width; x++ {
			oldPixel := arr.Array[x][y]

			arr.Array[x][y] = GetClosestColor(oldPixel)

			quantError := oldPixel - arr.Array[x][y]
			factorQuantError := float64(quantError) * factorErr

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
	}
	return arr
}

// Start position: Left Top
// Dither order: Left to right; continue with next row
// Matrix:
//  X  P  4  1
//  1  4  1
//     1
func DitheringMatrix3x4(a *common.Array, factorErr float64) *common.Array {
	arr := common.CloneArray(a)
	for y := arr.Height - 1; y >= 0; y-- {
		for x := 0; x < arr.Width; x++ {
			oldPixel := arr.Array[x][y]

			arr.Array[x][y] = GetClosestColor(oldPixel)

			quantError := oldPixel - arr.Array[x][y]
			factorQuantError := float64(quantError) * factorErr

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
	}
	return arr
}

// Start position: Left Top
// Dither order: Left to right; continue with next row
// Matrix:
//  X  X  P  8  4
//  2  4  8  4  2
//  1  2  4  2  1
func DitheringMatrix3x5(a *common.Array, factorErr float64) *common.Array {
	arr := common.CloneArray(a)
	for y := arr.Height - 1; y >= 0; y-- {
		for x := 0; x < arr.Width; x++ {
			oldPixel := arr.Array[x][y]

			arr.Array[x][y] = GetClosestColor(oldPixel)

			quantError := oldPixel - arr.Array[x][y]
			factorQuantError := float64(quantError) * factorErr

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
	return arr
}

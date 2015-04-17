package sequential

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image"
)

func RunSequentialMain(arr *common.Array, factorErr float64) *image.RGBA {
	arr = algorithm.DitheringMatrix2x3_2(arr, factorErr)
	//arr = algorithm.DitheringMatrix3x4(arr, 1.0)
	//arr = algorithm.DitheringMatrix3x5(arr, 1.0)

	dst := algorithm.ConvertGrayArrayToImage(arr)
	return dst
}

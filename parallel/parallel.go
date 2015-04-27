package parallel

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image"
	//"runtime"
)

var arr *common.Array

type Job struct {
	x, y, value int32
}

func RunParallelMain(ar *common.Array, factorErr float64) *image.RGBA {
	//cpus := runtime.NumCPU()
	arr = ar

	dst := algorithm.ConvertGrayArrayToImage(arr)
	return dst
}

func jobFactory(arr *common.Array, factorErr float64) chan Job {
	jobs := make(chan Job)

	return jobs

}

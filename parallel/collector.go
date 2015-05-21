package parallel

import (
	"strconv"

	"github.com/pichuio/halftone/common"
)

var WorkQueue = make(chan WorkRequest)
var GlobArray *common.Array

func Collector(arr *common.Array) {
	//do this only once and use the GlobalArray from this point
	GlobArray = common.CloneArray(arr)
	//for y := arr.Height - 1; y >= 0; y-- {
	//		for x := 0; x < arr.Width; x++ {
	//			oldPixel := arr.Array[x][y]
	//			s := string("worker_" + string(y) + "_" + string(x))
	//			work := WorkRequest{Name: s, X: y, Y: x, Value: oldPixel}

	//			WorkQueue <- work
	//
	//	}
	//}

	for x := 0; x < GlobArray.Width; x++ {
		oldPixel := GlobArray.Array[x][GlobArray.Height-1]
		s := "worker_" + strconv.Itoa(int(oldPixel))
		work := WorkRequest{Name: s, X: x, Y: GlobArray.Height - 1, Value: oldPixel}

		WorkQueue <- work
	}
}

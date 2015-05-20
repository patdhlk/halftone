package parallel

import (
	"strconv"

	"github.com/pichuio/halftone/common"
)

var WorkQueue = make(chan WorkRequest)

func Collector(arr *common.Array) {
	//for y := arr.Height - 1; y >= 0; y-- {
	//		for x := 0; x < arr.Width; x++ {
	//			oldPixel := arr.Array[x][y]
	//			s := string("worker_" + string(y) + "_" + string(x))
	//			work := WorkRequest{Name: s, X: y, Y: x, Value: oldPixel}

	//			WorkQueue <- work
	//
	//	}
	//}

	for x := 0; x < arr.Width; x++ {
		oldPixel := arr.Array[x][0]
		s := "worker_" + strconv.Itoa(int(oldPixel))
		work := WorkRequest{Name: s, X: x, Y: 0, Value: oldPixel}

		WorkQueue <- work
	}
}

package parallel

import (
	"image"
	"log"

	"github.com/pichuio/halftone/common"
)

func RunParallelMain(arr *common.Array, factorErr float64) *image.RGBA {
	log.Println("parallel started")
	StartDispatcher()
	Collector(arr)
	return nil
}

package parallel

import (
	//"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image"
	"log"
	"sync"
)

var arr *common.Array
var lock = &sync.Mutex{}

type job struct {
	line []int32 //use a slice when calling this, not an array
	y    int     //("vertical") line-number
}

type result struct {
	line []int32 //use a slice when calling this, not an array
	y    int     //("vertical") line-number
}

func jobFactory(ar *common.Array) chan job {
	jobs := make(chan job)
	go func() {
		for y := 0; y < ar.Height; y++ {
			j := new(job)
			j.line = make([]int32, ar.Width)
			for x := 0; x < ar.Width; x++ {
				j.line[x] = ar.Array[x][y]
				j.y = y
				jobs <- *j
			}
			log.Println(y, " ", j.line)
		}
		close(jobs)
	}()
	return jobs

}

func RunParallelMainMutex(ar *common.Array, factorErr float64) *image.RGBA {
	_ = jobFactory(ar)

	return nil
}

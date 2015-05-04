package parallel

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image"
	"log"
	"runtime"
	"sync"
)

var arr *common.Array
var mutex = &sync.Mutex{}

func RunParallelMainMutex(ar *common.Array, factorErr float64) *image.RGBA {
	//var k int
	progress := make([]int32, ar.Height)
	height := ar.Height
	width := ar.Width
	arr = ar
	output := common.CloneArray(ar)
	log.Println("CPUs: ", runtime.NumCPU())

	for k := 0; k < runtime.NumCPU(); k++ {
		go func() {
			for i := height - 1; i >= 0; i-- {

				mutex.Lock()

				for j := 0; j < width; j++ {

					log.Println(j, i)
					if arr.Array[j][i] < 128 {
						output.Array[j][i] = 0
					} else {
						output.Array[j][i] = 1
					}
					error := arr.Array[j][i] - 255*output.Array[j][i]

					if i+1 < width {
						arr.Array[j][i+1] += error * 7 / 16
					}
					if j+1 < height && i-1 > -1 {
						arr.Array[j+1][i-1] += error * 3 / 16
					}

					if j+1 < height {
						arr.Array[j+1][i] += error * 5 / 16
					}

					if j+1 < height && i+1 < width {
						arr.Array[j+1][i+1] += error * 1 / 16
					}

					progress[i] = int32(j + 1)

					if i == 2 && j+1 < height {
						mutex.Unlock()
					}
				}
			}
		}()
	}
	return algorithm.ConvertGrayArrayToImage(output)
}

package parallel

import (
	//"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image"
	"log"
	"sync"
)

var renderArray *common.Array //use this array for concurrent processing
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

func backgroundWorker(jobs <-chan job, results chan<- result, done chan<- bool) {
	for job := range jobs {
		results <- result{job.line, job.y}
	}
	done <- true
}

func resultCollector() {
	//TODO
}

func workerFactory(count int, jobs <-chan job, results chan<- result) {
	done := make(chan bool)
	for i := 0; i < count; i++ {
		go backgroundWorker(jobs, results, done)
	}

	go func() {
		for i := 0; i < count; i++ {
			<-done
		}
		close(results)
	}()
}

func RunParallelMainMutex(ar *common.Array, factorErr float64) *image.RGBA {
	//_ = jobFactory(ar)
	renderArray = ar
	return nil
}

package main

import (
	"runtime"
)

func main() {
	//USING ALL CORES OF YOUR MACHINE FOR PARALLEL PROCESSING
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)
}

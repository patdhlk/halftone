package main

import (
	"github.com/pichuio/halftone/common"
	"image/png"
	"log"
	"os"
	"runtime"
)

func main() {
	//USING ALL CORES OF YOUR MACHINE FOR PARALLEL PROCESSING
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	w := common.NewImageWorker()
	ic := common.NewImageConverter()
	img, err := w.LoadImage("Lenna.png")
	//img, err := w.LoadImage("testImage.png")

	if err != nil {
		log.Fatal(err)
	}

	width, height := w.GetImageDemensions(img)

	log.Println(width, height)

	if err != nil {
		log.Fatal(err)
	}

	gray := ic.ConvertToGray(img)

	outfilename := "result.png"
	outfile, err := os.Create(outfilename)
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()
	png.Encode(outfile, gray)

}

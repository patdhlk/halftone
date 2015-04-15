package main

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"image/color"
	"log"
	"runtime"
)

var pal = color.Palette{
	color.Black,
	color.White,
}

var DitherArray [][]int32

func main() {
	//USING ALL CORES OF YOUR MACHINE FOR PARALLEL PROCESSING
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	worker := common.NewImageWorker()
	ic := common.NewImageConverter()
	//img, err := worker.LoadImage("Lenna.png")
	img, err := worker.LoadImage("Michelangelo.png")
	//img, err := worker.LoadImage("sample.png")

	if err != nil {
		log.Fatal(err)
	}

	width, height := worker.GetImageDemensions(img)

	log.Println(width, height)

	if err != nil {
		log.Fatal(err)
	}

	grayImage := ic.ConvertToGray(img)

	worker.SaveImage("result.jpg", grayImage)

	log.Println("specify dest")

	arr := algorithm.ConvertImageToGrayArr(img)
	//gray array
	DitherArray = arr.Array
	//gray image
	dst := algorithm.ConvertGrayArrayToImage(arr)

	//save gray picture
	worker.SaveImage("result2.png", dst)

	//arr = algorithm.DitheringMatrix2x3_2(arr)
	//arr = algorithm.DitheringMatrix3x4(arr)
	arr = algorithm.DitheringMatrix3x5(arr)

	dst = algorithm.ConvertGrayArrayToImage(arr)

	//save gray picture
	worker.SaveImage("sequential.png", dst)
}

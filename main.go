package main

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"github.com/pichuio/halftone/sequential"
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

	log.Println("################## error diffusion ##################")
	log.Println("*****************************************************")
	log.Println("### supported image formats: *.jpg, *.jpeg, *.gif ###")
	log.Println("*****************************************************")
	log.Println("#####################################################")

	//USING ALL CORES OF YOUR MACHINE FOR PARALLEL PROCESSING
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	worker := common.NewImageWorker()
	ic := common.NewImageConverter()
	//img, err := worker.LoadImage("images/original/Lenna.png")
	img, err := worker.LoadImage("images/original/Michelangelo.png")

	if err != nil {
		log.Fatal(err)
	}

	width, height := worker.GetImageDemensions(img)

	log.Println(width, height)

	if err != nil {
		log.Fatal(err)
	}

	grayImage := ic.ConvertToGray(img)

	worker.SaveImage("images/processing/seq_grayImage.png", grayImage)

	arr := algorithm.ConvertImageToGrayArr(img)
	//gray array
	DitherArray = arr.Array
	//gray image
	dst := algorithm.ConvertGrayArrayToImage(arr)
	//save gray picture
	worker.SaveImage("images/processing/seq_grayImage2.png", dst)

	//sequential processing
	dst = sequential.RunSequentialMain(arr, 0.9)

	//parallel processing
	//TODO

	//save gray picture
	worker.SaveImage("images/processing/seq_result.png", dst)
}

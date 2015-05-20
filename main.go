package main

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"github.com/pichuio/halftone/parallel"
	"github.com/pichuio/halftone/sequential"
	"log"
	"runtime"
)

func main() {

	log.Println("################## error diffusion ##################")
	log.Println("*****************************************************")
	log.Println("### supported image formats: *.jpg, *.jpeg, *.gif ###")
	log.Println("*****************************************************")
	log.Println("#####################################################")

	//USING ALL CORES OF YOUR MACHINE FOR PARALLEL PROCESSING
	numcpu := runtime.NumCPU()
	log.Println(numcpu)
	runtime.GOMAXPROCS(numcpu)

	//create processing directory, because if dir not exists the outputimages cannot found by the user
	go common.CreateDirIfNotExist("images/processing")

	worker := common.NewImageWorker()
	ic := common.NewImageConverter()
	//img, err := worker.LoadImage("images/original/Lenna.png")
	img, err := worker.LoadImage("images/original/Michelangelo.png")
	//img, err := worker.LoadImage("images/original/sample.jpg")

	if err != nil {
		log.Fatal(err)
	}

	//
	width, height := worker.GetImageDemensions(img)

	log.Println(width, height)

	if err != nil {
		log.Fatal(err)
	}

	grayImage := ic.ConvertToGray(img)

	worker.SaveImage("images/processing/seq_grayImage.png", grayImage)

	arr := algorithm.ConvertImageToGrayArr(img)
	//gray image
	dst := algorithm.ConvertGrayArrayToImage(arr)
	//save gray picture
	worker.SaveImage("images/processing/seq_grayImage2.png", dst)

	log.Println("start dithering")
	//sequential processing
	dst = sequential.RunSequentialMain(arr, 0.9)

	//parallel processing
	dst_par := parallel.RunParallelMain(arr, 0.9) //TODO

	//save gray picture
	worker.SaveImage("images/processing/seq_result.png", dst)
	if dst_par != nil {
		worker.SaveImage("images/processing/par_result.png", dst_par)
	}
	log.Println("FINISHED")
}

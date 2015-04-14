package main

import (
	"github.com/pichuio/halftone/common"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime"
)

var pal = color.Palette{
	color.Black,
	color.White,
}

var DitherArray [][]uint32

func main() {
	//USING ALL CORES OF YOUR MACHINE FOR PARALLEL PROCESSING
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	worker := common.NewImageWorker()
	ic := common.NewImageConverter()
	img, err := worker.LoadImage("Lenna.png")
	//img, err := w.LoadImage("testImage.png")

	if err != nil {
		log.Fatal(err)
	}

	width, height := worker.GetImageDemensions(img)

	log.Println(width, height)

	if err != nil {
		log.Fatal(err)
	}

	grayImage := ic.ConvertToGray(img)

	worker.SaveImage("result.png", grayImage)

	//the image bounds
	bounds := grayImage.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	log.Println("specify dest")

	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	//create 2D array

	DitherArray = make([][]uint32, w)
	for i := 0; i < w; i++ {
		DitherArray[i] = make([]uint32, h)
	}

	//http://dotnet-snippets.de/snippet/floyd-steinberg-dithering/94
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//use colored image
			pixel := img.At(y, x)
			//use gray image
			//pixel := img.At(y, x)
			red, green, blue, _ := pixel.RGBA()
			gray := CalculateGray(red, green, blue)

			DitherArray[x][y] = gray
		}
	}

	for x := 1; x < h-1; x++ {
		for y := 1; y < w-1; y++ {
			CalculateDithering(x, y)
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//pixel := grayImage.At(y, x)
			temp := DitherArray[x][y]
			if temp == 0 {
				temp = 0
			} else {
				temp = 255
			}
			tem := uint8(temp)
			log.Println(tem)
			c := color.RGBA{tem, tem, tem, 0xff}
			dst.SetRGBA(x, y, c)
		}
	}

	worker.SaveImage("sequential.png", dst)
}

func CalculateGray(red, green, blue uint32) uint32 {
	return (red + green + blue) / 3
}

func CalculateDithering(line, column int) {
	var factor int

	if DitherArray[line][column] < 128 {
		factor = int(DitherArray[line][column] / 16)
		DitherArray[line][column] = 0
	} else {
		factor = int((DitherArray[line][column] - 255) / 16)
		DitherArray[line][column] = 1
	}

	DitherArray[line+1][column-1] += uint32(factor * 3)
	DitherArray[line+1][column] += uint32(factor * 5)
	DitherArray[line+1][column+1] += uint32(factor)
	DitherArray[line][column+1] += uint32(factor * 7)
}

func saveImage(path string, i image.Image) {
	w, _ := os.Create(path)
	if err := png.Encode(w, i); err != nil {
		log.Println("Error writing image on disk")
		os.Exit(1)
	}
}

package main

import (
	"github.com/pichuio/halftone/common"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"runtime"
)

var pal = color.Palette{
	color.Black,
	color.White,
}

var DitherArray [][]uint8

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

	DitherArray = make([][]uint8, w)
	for i := 0; i < w; i++ {
		DitherArray[i] = make([]uint8, h)
	}

	//http://dotnet-snippets.de/snippet/floyd-steinberg-dithering/94
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//use colored image
			pixel := img.At(x, y)
			//use gray image
			//pixel := img.At(y, x)
			red, green, blue, _ := pixel.RGBA()
			red2 := uint8(red)
			green2 := uint8(green)
			blue2 := uint8(blue)
			gray := CalculateGray(red2, green2, blue2)
			gray_ := uint8(gray)
			DitherArray[x][y] = gray_
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//pixel := grayImage.At(y, x)
			temp := DitherArray[x][y]

			c := color.RGBA{temp, temp, temp, 0xff}
			dst.SetRGBA(x, y, c)
		}
	}

	worker.SaveImage("result2.png", dst)

	for x := 1; x < w-1; x++ {
		for y := 1; y < h-1; y++ {
			CalculateDithering(x, y)
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			//pixel := grayImage.At(y, x)
			temp := DitherArray[x][y]
			if temp == 0 {
				temp = math.MaxUint8
			} else {
				temp = 0
			}

			c := color.RGBA{temp, temp, temp, 0xff}
			dst.SetRGBA(x, y, c)
		}
	}

	worker.SaveImage("sequential.png", dst)
}

func CalculateGray(red, green, blue uint8) uint32 {
	return (uint32(red) + uint32(green) + uint32(blue)) / 3
}

func CalculateDithering(x, y int) {
	var factor int

	var act uint8 = uint8(DitherArray[x][y])

	if act < 128 {
		factor = int(act / 16)
		DitherArray[x][y] = 0
	} else {
		factor = int((act - 255) / 16)
		DitherArray[x][y] = 1
	}

	DitherArray[x+1][y-1] += uint8(factor * 3)
	DitherArray[x+1][y] += uint8(factor * 5)
	DitherArray[x+1][y+1] += uint8(factor)
	DitherArray[x][y+1] += uint8(factor * 7)
}

func saveImage(path string, i image.Image) {
	w, _ := os.Create(path)
	if err := png.Encode(w, i); err != nil {
		log.Println("Error writing image on disk")
		os.Exit(1)
	}
}

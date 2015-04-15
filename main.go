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

	worker.SaveImage("result.png", grayImage)

	//the image bounds
	bounds := grayImage.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	log.Println("specify dest")

	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	//create 2D array

	DitherArray = make([][]int32, w)
	for i := 0; i < w; i++ {
		DitherArray[i] = make([]int32, h)
	}

	//http://dotnet-snippets.de/snippet/floyd-steinberg-dithering/94
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//use colored image
			//pixel := grayImage.At(x, y)
			//use gray image
			pixel := img.At(x, y)
			red, green, blue, _ := pixel.RGBA()
			red2 := uint8(red)
			green2 := uint8(green)
			blue2 := uint8(blue)
			gray := CalculateGray(red2, green2, blue2)
			gray_ := uint8(gray)
			DitherArray[x][y] = int32(gray_)
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//pixel := grayImage.At(y, x)
			temp := DitherArray[x][y]

			c := color.RGBA{uint8(temp), uint8(temp), uint8(temp), 0xff}
			dst.SetRGBA(x, y, c)
		}
	}

	worker.SaveImage("result2.png", dst)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			CalculateDithering(x, y, w, h)
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			//pixel := grayImage.At(y, x)
			temp := DitherArray[x][y]
			/*
				if temp == 0 {
					temp = math.MaxUint8
				} else {
					temp = 0
				}
			*/

			c := color.RGBA{uint8(temp), uint8(temp), uint8(temp), 0xff}
			dst.SetRGBA(x, y, c)
		}
	}

	worker.SaveImage("sequential.png", dst)
}

func CalculateGray(red, green, blue uint8) uint32 {
	return (uint32(red) + uint32(green) + uint32(blue)) / 3
}

func CalculateDithering(x, y, w, h int) {
	oldPixel := DitherArray[x][y]

	if oldPixel < 128 {
		DitherArray[x][y] = 0
	} else {
		DitherArray[x][y] = 255
	}

	quantError := oldPixel - DitherArray[x][y]

	if y-1 >= 0 && x+1 < w {
		DitherArray[x+1][y-1] += int32(float32(quantError) * (3.0 / 16.0))
	}
	if x+1 < w {
		DitherArray[x+1][y] += int32(float32(quantError) * (5.0 / 16.0))
	}
	if y+1 < h && x+1 < w {
		DitherArray[x+1][y+1] += int32(float32(quantError) * (1.0 / 16.0))
	}
	if y+1 < h {
		DitherArray[x][y+1] += int32(float32(quantError) * (7.0 / 16.0))
	}
}

func saveImage(path string, i image.Image) {
	w, _ := os.Create(path)
	if err := png.Encode(w, i); err != nil {
		log.Println("Error writing image on disk")
		os.Exit(1)
	}
}

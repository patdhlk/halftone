package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/patdhlk/halftone"
)

var (
	rng           = rand.New(rand.NewSource(time.Now().UnixNano()))
	ditherModeVar int
)

func init() {
	flag.IntVar(&ditherModeVar, "mode", 0, "specifies the dither mode")
}

func main() {
	flag.Parse()
	files := []string{"img/Lenna.png", "img/Michelangelo.png", "img/radon.jpg", "img/sample.jpg", "img/timon.jpg"}
	worker := halftone.NewImageWorker()
	cv := halftone.NewImageConverter()
	for _, file := range files {
		var img, err = worker.LoadImage(file)

		if err != nil {
			log.Fatal(err)
		}

		var gray = cv.ConvertToGray(img)
		var dithered *image.Gray
		switch ditherModeVar {
		case 0:
			dithered = halftone.FloydSteinbergDitherer{}.Run(gray)
		case 1:
			dithered = halftone.NewGridDitherer(5, 3, 8, rng).Run(gray)
		case 2:
			dithered = halftone.NewThresholdDitherer(122).Run(gray)
		default:
			fmt.Println("wrong dither mode specified. Only 0, 1 and 2 supported")
			os.Exit(1)
		}

		// Save as out.png
		newFilename := strings.Replace(file, "img/", "out/", 1)
		f, err := os.Create(newFilename)
		if err != nil {
			fmt.Printf("%s", err.Error())
			os.Exit(1)
		}
		defer f.Close()
		png.Encode(f, dithered)
	}
}

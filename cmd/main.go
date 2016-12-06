package main

import (
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"
    "log"

	"github.com/patdhlk/halftone"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	files := []string{"img/Lenna.png", "img/Michelangelo.png", "img/radon.jpg", "img/sample.jpg", "img/timon.jpg"}
	worker := halftone.NewImageWorker()
    cv := halftone.NewImageConverter() 
	for _, file := range files {
		var img, err = worker.LoadImage(file)

        if err != nil {
            log.Fatal(err) 
        }

		var gray = cv.ConvertToGray(img)
		//var dithered = ThresholdDitherer{122}.apply(gray)
		//var dithered = GridDitherer{5, 3, 8, rng}.apply(gray)
		var dithered = halftone.FloydSteinbergDitherer{}.Run(gray)

		// Save as out.png
		newFilename := strings.Replace(file, "img/", "out/", 1)
		f, _ := os.Create(newFilename)
		defer f.Close()
		png.Encode(f, dithered)
	}
}

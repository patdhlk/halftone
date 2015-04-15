package main

import (
	"github.com/pichuio/halftone/common"
	"testing"
)

func TestDitherResult(t *testing.T) {
	worker := common.NewImageWorker()
	img1, _ := worker.LoadImage("sequential.png")
	img2, _ := worker.LoadImage("Michelangelo_FloydWiki.png")

	width1, height1 := worker.GetImageDemensions(img1)

	width2, height2 := worker.GetImageDemensions(img2)

	if width1 != width2 {
		t.Errorf("TestDitherResult - Width %v %v", width1, width2)
	}

	if height1 != height2 {
		t.Errorf("TestDitherResult - Height %v %v", height1, height2)
	}

	errCount := 0

	for x := 0; x < width1; x++ {
		for y := 0; y < height1; y++ {
			pixel := img1.At(x, y)
			red1, green1, blue1, _ := pixel.RGBA()

			pixel = img2.At(x, y)
			red2, green2, blue2, _ := pixel.RGBA()

			if red1 != red2 || green1 != green2 || blue1 != blue2 {
				//t.Errorf("TestDitherResult - Red %v %v Green %v %v Blue %v %v", red1, red2, green1, green2, blue1, blue2)
				errCount++
			}
		}
	}
	if errCount > 0 {
		t.Errorf("TestDitherResult - Count %v (%v)", errCount, height1*width1)
	}
}

func TestColorToGray(t *testing.T) {
	worker := common.NewImageWorker()
	img1, _ := worker.LoadImage("result.png")
	img2, _ := worker.LoadImage("result2.png")

	width1, height1 := worker.GetImageDemensions(img1)

	width2, height2 := worker.GetImageDemensions(img2)

	if width1 != width2 {
		t.Errorf("TestDitherResult - Width %v %v", width1, width2)
	}

	if height1 != height2 {
		t.Errorf("TestDitherResult - Height %v %v", height1, height2)
	}

	for x := 0; x < width1; x++ {
		for y := 0; y < height1; y++ {
			pixel := img1.At(x, y)
			red1, green1, blue1, _ := pixel.RGBA()

			pixel = img2.At(x, y)
			red2, green2, blue2, _ := pixel.RGBA()

			if red1 != red2 {
				t.Errorf("TestDitherResult - Red %v %v", red1, red2)
			}

			if green1 != green2 {
				t.Errorf("TestDitherResult - Green %v %v", green1, green2)
			}

			if blue1 != blue2 {
				t.Errorf("TestDitherResult - Blue %v %v", blue1, blue2)
			}
		}
	}
}

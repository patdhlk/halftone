package main

import (
	"github.com/pichuio/halftone/algorithm"
	"github.com/pichuio/halftone/common"
	"testing"
)

func TestDitherResult(t *testing.T) {
	//	worker := common.NewImageWorker()
	//	img1, _ := worker.LoadImage("images/processing/seq_result.png")
	//	img2, _ := worker.LoadImage("images/original/Michelangelo_Result.png")

	//	width1, height1 := worker.GetImageDemensions(img1)

	//	width2, height2 := worker.GetImageDemensions(img2)

	//	if width1 != width2 {
	//		t.Errorf("TestDitherResult - Width %v %v", width1, width2)
	//	}

	//	if height1 != height2 {
	//		t.Errorf("TestDitherResult - Height %v %v", height1, height2)
	//	}

	//	errCount := 0

	//	for x := 0; x < width1; x++ {
	//		for y := 0; y < height1; y++ {
	//			pixel := img1.At(x, y)
	//			red1, green1, blue1, _ := pixel.RGBA()

	//			pixel = img2.At(x, y)
	//			red2, green2, blue2, _ := pixel.RGBA()

	//			if red1 != red2 || green1 != green2 || blue1 != blue2 {
	//				//t.Errorf("TestDitherResult - Red %v %v Green %v %v Blue %v %v", red1, red2, green1, green2, blue1, blue2)
	//				errCount++
	//			}
	//		}
	//	}
	//	if errCount > 0 {
	//		t.Errorf("TestDitherResult - Count %v (%v)", errCount, height1*width1)
	//	}
}

func TestColorToGray(t *testing.T) {
	worker := common.NewImageWorker()
	img1, _ := worker.LoadImage("images/processing/seq_grayImage.png")
	img2, _ := worker.LoadImage("images/processing/seq_grayImage2.png")

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

func TestDitherResult2(t *testing.T) {
	grayArr := common.NewArray(200, 100)
	for x := 0; x < grayArr.Width; x++ {
		for y := 0; y < grayArr.Height; y++ {
			grayArr.Array[x][y] = 127
		}
	}
	newArr := algorithm.DitheringMatrix2x3_2(grayArr, 1.0)
	for x := 0; x < grayArr.Width; x++ {
		for y := 0; y < grayArr.Height; y++ {
			xMod2 := x%2 == 0
			yMod2 := y%2 == 0

			if (xMod2 && !yMod2) || (!xMod2 && yMod2) {
				//BLACK
				if newArr.Array[x][newArr.Height-1-y] != 255 {
					t.Errorf("TestDitherResult2 != 255 - [%v][%v] %v", x, newArr.Height-1-y, newArr.Array[x][newArr.Height-1-y])
				}
			} else {
				if newArr.Array[x][newArr.Height-1-y] != 0 {
					t.Errorf("TestDitherResult2 != 0 - [%v][%v] %v", x, newArr.Height-1-y, newArr.Array[x][newArr.Height-1-y])
				}
			}
		}
	}

	dst := algorithm.ConvertGrayArrayToImage(newArr)

	worker := common.NewImageWorker()
	worker.SaveImage("images/testing/chessboard.png", dst)
}

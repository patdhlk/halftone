package sequential

import (
	"github.com/pichuio/halftone/common"
	"testing"
)

func TestDitherResultGrayImage(t *testing.T) {
	grayArr := common.NewArray(200, 100)
	for x := 0; x < grayArr.Width; x++ {
		for y := 0; y < grayArr.Height; y++ {
			grayArr.Array[x][y] = 127
		}
	}

	newArr0 := Dithering(grayArr, 0, 1.0)
	newArr1 := Dithering(grayArr, 0, 1.0)
	newArr2 := Dithering(grayArr, 0, 1.0)
	for x := 0; x < grayArr.Width; x++ {
		for y := 0; y < grayArr.Height; y++ {
			xMod2 := x%2 == 0
			yMod2 := y%2 == 0

			if (xMod2 && !yMod2) || (!xMod2 && yMod2) {
				//BLACK
				if newArr0.Array[x][newArr0.Height-1-y] != 255 {
					t.Errorf("TestDitherResult2 != 255 - [%v][%v] %v", x, newArr0.Height-1-y, newArr0.Array[x][newArr0.Height-1-y])
				}
				if newArr1.Array[x][newArr1.Height-1-y] != 255 {
					t.Errorf("TestDitherResult2 != 255 - [%v][%v] %v", x, newArr1.Height-1-y, newArr1.Array[x][newArr1.Height-1-y])
				}
				if newArr2.Array[x][newArr2.Height-1-y] != 255 {
					t.Errorf("TestDitherResult2 != 255 - [%v][%v] %v", x, newArr2.Height-1-y, newArr2.Array[x][newArr2.Height-1-y])
				}
			} else {
				if newArr0.Array[x][newArr0.Height-1-y] != 0 {
					t.Errorf("TestDitherResult2_0 != 0 - [%v][%v] %v", x, newArr0.Height-1-y, newArr0.Array[x][newArr0.Height-1-y])
				}
				if newArr1.Array[x][newArr1.Height-1-y] != 0 {
					t.Errorf("TestDitherResult2_1 != 0 - [%v][%v] %v", x, newArr1.Height-1-y, newArr1.Array[x][newArr1.Height-1-y])
				}
				if newArr2.Array[x][newArr0.Height-1-y] != 0 {
					t.Errorf("TestDitherResult2_2 != 0 - [%v][%v] %v", x, newArr2.Height-1-y, newArr2.Array[x][newArr2.Height-1-y])
				}
			}
		}
	}
}

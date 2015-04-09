package common

import (
	"image"
	"image/png"
	"log"
	"os"
)

type ImageWorker struct {
}

func NewImageWorker() *ImageWorker {
	w := new(ImageWorker)
	return w
}

func (w *ImageWorker) LoadImage(path string) (image.Image, error) {

	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()
	img, err := png.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	return img, err
}

func (w *ImageWorker) GetImageDemensions(image image.Image) (int, int) {
	bounds := image.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	return width, height
}

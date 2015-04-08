package common

import (
	"image"
	"os"
)

type ImageLoader struct {
}

func NewImageLoader() *ImageLoader {
	il := new(ImageLoader)
	return il
}

func (il *ImageLoader) LoadImage(path string) (image.Image, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

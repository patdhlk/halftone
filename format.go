package halftone

import (
	"errors"
)

type Format int

const (
	JPEG Format = iota
	PNG
	GIF
	TIFF
	BMP
)

func (f Format) String() string {
	switch f {
	case JPEG:
		return "JPEG"
	case PNG:
		return "PNG"
	case GIF:
		return "GIF"
	/*case TIFF:
		return "TIFF"
	case BMP:
		return "BMP"*/
	default:
		return "Unsupported"
	}
}

var (
	ErrUnsupportedFormat = errors.New("imaging: unsupported image format")
)

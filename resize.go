package eggbite

import (
	"image"

	"github.com/nfnt/resize"
)

func ResizeImageBytes(imgBytes []byte, width, height uint) (image.Image, error) {
	img, e := ImageFromBytes(imgBytes)
	if e != nil {
		return nil, e
	}
	resizedImg := resize.Resize(width, height, img, resize.Bilinear)
	return resizedImg, nil
}

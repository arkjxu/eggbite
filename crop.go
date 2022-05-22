package eggbite

import (
	"fmt"
	"image"
	"image/draw"
)

func CropBytes(imgBytes []byte, x1, y1, x2, y2 int) (image.Image, error) {
	img, e := ImageFromBytes(imgBytes)
	if e != nil {
		return nil, e
	}

	imgType := fmt.Sprintf("%T", img)
	if imgType == "*image.NRGBA" {
		imgBounds := img.Bounds()
		tmpImg := image.NewRGBA(image.Rect(0, 0, imgBounds.Dx(), imgBounds.Dy()))
		draw.Draw(tmpImg, tmpImg.Bounds(), img, imgBounds.Min, draw.Src)
		img = tmpImg
	}

	croppedImage := img.(*image.RGBA).SubImage(image.Rect(x1, y1, x2, y2))
	return croppedImage, e
}

func CropFile(filename string, x1, y1, x2, y2 int) (image.Image, error) {
	img, e := ImageFromFile(filename)
	if e != nil {
		return nil, e
	}
	croppedImage := img.(*image.RGBA).SubImage(image.Rect(x1, y1, x2, y2))
	return croppedImage, e
}

func CropImage(img image.Image, x1, y1, x2, y2 int) (image.Image, error) {
	imgType := fmt.Sprintf("%T", img)
	if imgType == "*image.NRGBA" {
		imgBounds := img.Bounds()
		tmpImg := image.NewRGBA(image.Rect(0, 0, imgBounds.Dx(), imgBounds.Dy()))
		draw.Draw(tmpImg, tmpImg.Bounds(), img, imgBounds.Min, draw.Src)
		img = tmpImg
	}
	croppedImage := img.(*image.RGBA).SubImage(image.Rect(x1, y1, x2, y2))
	return croppedImage, nil
}

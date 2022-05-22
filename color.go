package eggbite

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
)

var (
	ErrInvalidFormat = errors.New("invalid color format")
)

func IsColor(color1, color2 color.Color, similarity float64) (isColor bool) {
	r1, g1, b1 := withMaxRGBValue(color1.RGBA())
	r2, g2, b2 := withMaxRGBValue(color2.RGBA())

	rDiff := math.Abs(float64(r1 - r2))
	gDiff := math.Abs(float64(g1 - g2))
	bDiff := math.Abs(float64(b1 - b2))

	if 1-(rDiff+gDiff+bDiff)/255/3 >= similarity {
		isColor = true
	}
	return isColor
}

func IsColorWithOffset(color1, color2, offset color.Color, similarity float64) bool {
	rOff, gOff, bOff := withMaxRGBValue(offset.RGBA())
	r2, g2, b2 := withMaxRGBValue(color2.RGBA())

	rMax := addColorValue(r2, rOff)
	gMax := addColorValue(g2, gOff)
	bMax := addColorValue(b2, bOff)
	rFind := subColorValue(r2, rOff)

	for {
		gFind := subColorValue(g2, gOff)
		for {
			bFind := subColorValue(b2, bOff)
			for {
				offColor := color.RGBA{R: rFind, G: gFind, B: bFind}
				isColorFound := IsColor(color1, offColor, similarity)
				if isColorFound {
					return true
				}
				if bFind == bMax {
					break
				}
				bFind = addColorValue(bFind, 1)
			}
			if gFind == gMax {
				break
			}
			gFind = addColorValue(gFind, 1)
		}
		if rFind == rMax {
			break
		}
		rFind = addColorValue(rFind, 1)
	}
	return false
}

func FindColor(src image.Image, colorToFind color.Color, similarity float64) (x, y int) {
	x, y = -1, -1

	imageSize := src.Bounds().Size()

	width := imageSize.X
	height := imageSize.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colorAtPix := src.At(x, y)
			isColorFound := IsColor(colorAtPix, colorToFind, similarity)
			if isColorFound {
				return x, y
			}
		}
	}
	return x, y
}

func FindColorWithOffset(src image.Image, colorToFind color.Color, offset color.Color, similarity float64) (x, y int) {
	x, y = -1, -1

	imageSize := src.Bounds().Size()

	width := imageSize.X
	height := imageSize.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colorAtPix := src.At(x, y)
			isColorFound := IsColorWithOffset(colorAtPix, colorToFind, offset, similarity)
			if isColorFound {
				return x, y
			}
		}
	}

	return x, y
}

func GetColorNum(img image.Image, colorToFind color.Color, similarity float64) (uint, error) {
	imageSize := img.Bounds().Size()

	width := imageSize.X
	height := imageSize.Y
	colorCount := uint(0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colorAtPix := img.At(x, y)
			isColorFound := IsColor(colorAtPix, colorToFind, similarity)
			if isColorFound {
				colorCount += 1
			}
		}
	}
	return colorCount, nil
}

func GetColorNumWithOffset(img image.Image, colorToFind color.Color, offset color.Color, similarity float64) (uint, error) {
	imageSize := img.Bounds().Size()

	width := imageSize.X
	height := imageSize.Y

	colorCount := uint(0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colorAtPix := img.At(x, y)
			isColorFound := IsColorWithOffset(colorAtPix, colorToFind, offset, similarity)
			if isColorFound {
				colorCount += 1
			}
		}
	}

	return colorCount, nil
}

func RGBToHex(src color.Color) string {
	r, g, b, _ := src.RGBA()
	return fmt.Sprintf("%02X%02X%02X", uint8(r), uint8(g), uint8(b))
}

func HexToRGB(hex string) (c color.RGBA, e error) {
	c.A = 0xff

	if hex[0] != '#' {
		return c, ErrInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		e = ErrInvalidFormat
		return 0
	}

	switch len(hex) {
	case 7:
		c.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		c.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		c.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		c.R = hexToByte(hex[1]) * 17
		c.G = hexToByte(hex[2]) * 17
		c.B = hexToByte(hex[3]) * 17
	default:
		e = ErrInvalidFormat
	}
	c.A = 255
	return c, e
}

func withMaxRGBValue(r, g, b, _ uint32) (uint8, uint8, uint8) {
	return uint8(r & 255), uint8(g & 255), uint8(b & 255)
}

func addColorValue(c1, c2 uint8) uint8 {
	c1U := uint16(c1)
	c2U := uint16(c2)
	result := c1U + c2U
	if result > 255 {
		return 255
	}
	return uint8(result)
}

func subColorValue(c1, c2 uint8) uint8 {
	c1U := int16(c1)
	c2U := int16(c2)
	result := c1U - c2U
	if result < 0 {
		return 0
	}
	return uint8(result)
}

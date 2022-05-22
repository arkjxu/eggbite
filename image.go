package eggbite

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/corona10/goimagehash"
)

type HashType func(img image.Image) (*goimagehash.ImageHash, error)

var (
	Perception HashType = goimagehash.PerceptionHash
	Difference HashType = goimagehash.DifferenceHash
)

func FindImage(src image.Image, toFind image.Image) (x, y int) {
	x, y = -1, -1

	return x, y
}

func GetImageSimilarity(src image.Image, compareWith image.Image, ht HashType) (float64, error) {
	srcHash, e := ht(src)
	if e != nil {
		return 0, e
	}
	compareWithHash, e := ht(compareWith)
	if e != nil {
		return 0, e
	}
	return percentageDifference(srcHash.GetHash(), compareWithHash.GetHash()), nil
}

func ImageFromBytes(imgBytes []byte) (image.Image, error) {
	img, _, e := image.Decode(bytes.NewReader(imgBytes))
	if e != nil {
		return nil, e
	}
	return img, nil
}

func ImageFromFile(filename string) (image.Image, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func percentageDifference(srcHash, destHash uint64) float64 {
	binaryRepresentation := strconv.FormatInt(int64(srcHash^destHash), 2)
	percentageDiff := math.Round((float64((64-strings.Count(binaryRepresentation, "1"))*100.0)/64.0)*10000) / 10000
	return percentageDiff
}

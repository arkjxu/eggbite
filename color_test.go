package eggbite_test

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/arkjxu/eggbite"
)

func TestConvertRGBToHex(t *testing.T) {
	testColorHex := "F84119"
	testColor := color.RGBA{R: 248, G: 65, B: 25}
	convertedHex := eggbite.RGBToHex(testColor)
	if convertedHex != testColorHex {
		t.Fatal(fmt.Errorf("expecting %s but got %s", testColorHex, convertedHex))
	}
}

func TestConvertHexToRGB(t *testing.T) {
	testColorHex := "#F84119"
	testColor := color.RGBA{R: 248, G: 65, B: 25}
	convertedColor, e := eggbite.HexToRGB(testColorHex)
	if e != nil {
		t.Fatal(e)
	}
	rExpected, gExpect, bExpected, _ := testColor.RGBA()
	rActual, gActual, bActual, _ := convertedColor.RGBA()
	if rExpected != rActual || gExpect != gActual || bExpected != bActual {
		t.Fatal(e)
	}
}

func TestIsColor(t *testing.T) {
	testColor1 := color.RGBA{
		R: 255,
		G: 87,
		B: 51,
	}
	testColor2 := color.RGBA{
		R: 255,
		G: 87,
		B: 52,
	}
	shouldNotBeSame := eggbite.IsColor(testColor1, testColor2, 1)
	if shouldNotBeSame {
		t.Fatal(fmt.Errorf("color should not be the same"))
	}
	shouldBeSame := eggbite.IsColor(testColor1, testColor1, 1)
	if !shouldBeSame {
		t.Fatal(fmt.Errorf("color should be the same"))
	}
}

func TestIsColorWithOffset(t *testing.T) {
	testColor1 := color.RGBA{
		R: 255,
		G: 87,
		B: 51,
	}
	testColor2 := color.RGBA{
		R: 255,
		G: 88,
		B: 52,
	}
	offsetColor := color.RGBA{
		R: 161,
		G: 1,
		B: 1,
	}
	shouldBeFound := eggbite.IsColorWithOffset(testColor1, testColor2, offsetColor, 1)
	if !shouldBeFound {
		t.Fatal(fmt.Errorf("color with offset should be the same"))
	}
	testColor2.G = 89
	shouldBeFound = eggbite.IsColorWithOffset(testColor1, testColor2, offsetColor, 1)
	if shouldBeFound {
		t.Fatal(fmt.Errorf("color with offset should not be the same"))
	}
}

func TestGetColorNum(t *testing.T) {
	expectedFound := uint(16)
	fs, e := os.Open("images/reset-code.png")
	if e != nil {
		t.Fatal(e)
	}
	defer fs.Close()
	img, e := png.Decode(fs)
	if e != nil {
		t.Fatal(e)
	}
	count, e := eggbite.GetColorNum(img, color.RGBA{R: 147, G: 130, B: 88}, 1)
	if e != nil {
		t.Fatal(e)
	}
	if count != expectedFound {
		t.Fatalf("should only find %d, but found %d", expectedFound, count)
	}
}

func TestGetColorNumWithOffset(t *testing.T) {
	expectedFound := uint(16)
	fs, e := os.Open("images/reset-code.png")
	if e != nil {
		t.Fatal(e)
	}
	defer fs.Close()
	img, e := png.Decode(fs)
	if e != nil {
		t.Fatal(e)
	}
	count, e := eggbite.GetColorNumWithOffset(img, color.RGBA{R: 147, G: 130, B: 81}, color.RGBA{R: 0, G: 0, B: 7}, 1)
	if e != nil {
		t.Fatal(e)
	}
	if count != expectedFound {
		t.Fatalf("should only find %d, but found %d", expectedFound, count)
	}
}

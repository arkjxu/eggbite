package eggbite

import (
	"errors"
	"path"
	"runtime"

	"github.com/otiai10/gosseract/v2"
)

type OCRLanguage string

const (
	English           OCRLanguage = "eng"
	ChineseSimplified OCRLanguage = "chi_sim"
)

var (
	_eggbiteOcr *gosseract.Client
)

func InitOCR(language OCRLanguage) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("unable to get runtime location"))
	}
	tessDataPath := path.Join(path.Dir(filename), "tessdata")

	_eggbiteOcr = gosseract.NewClient()
	_eggbiteOcr.SetTessdataPrefix(tessDataPath)
	_eggbiteOcr.SetLanguage(string(language))
}

func FindTextWithBytes(imgBytes []byte) (string, error) {
	e := _eggbiteOcr.SetImageFromBytes(imgBytes)
	if e != nil {
		return "", e
	}
	text, e := _eggbiteOcr.Text()
	if e != nil {
		return "", e
	}
	return text, nil
}

func FindTextWithFile(fileName string) (string, error) {
	e := _eggbiteOcr.SetImage(fileName)
	if e != nil {
		return "", e
	}
	text, e := _eggbiteOcr.Text()
	if e != nil {
		return "", e
	}
	return text, nil
}

func CloseOCR() {
	defer _eggbiteOcr.Close()
}

package eggbite_test

import (
	"testing"

	"github.com/arkjxu/eggbite"
)

func TestOCRNumberCode(t *testing.T) {
	eggbite.InitOCR(eggbite.English)
	defer eggbite.CloseOCR()
	text, e := eggbite.FindTextWithFile("images/reset-code.png")
	if e != nil {
		t.Fatal(e)
	}
	if text != "406340" {
		t.Fatalf("expected 406340, but got %s\n", text)
	}
}

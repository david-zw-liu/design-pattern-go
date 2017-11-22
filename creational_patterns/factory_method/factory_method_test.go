package factory_method

import (
	"strings"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var di *DecodedImage
	var ir ImageReader
	image := "image.jpg"

	if strings.HasSuffix(image, ".jpg") {
		ir = CreateJpgeReader(image)
	} else if strings.HasSuffix(image, ".gif") {
		ir = CreateGifReader(image)
	}

	if ir == nil {
		t.Error("ir expected to get ImageReader (factory), but got nil")
	}

	di = ir.GetDecodeImage()
	if di.decodedImage != "jpge decoded: image.jpg" {
		t.Errorf("di.decodedImage expected to get \"jpge decoded: image.jpg\", but got %v", di.decodedImage)
	}

}

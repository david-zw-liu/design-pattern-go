package factory_method

import "fmt"

type DecodedImage struct {
	decodedImage string
}

func CreateDecodedImage(image string) *DecodedImage {
	return &DecodedImage{decodedImage: image}
}

type ImageReader interface {
	GetDecodeImage() *DecodedImage
}

type GifReader struct {
	decodedImg *DecodedImage
}

func CreateGifReader(image string) *GifReader {
	decodedImage := fmt.Sprintf("gif decoded: %s", image)
	return &GifReader{decodedImg: CreateDecodedImage(decodedImage)}
}

func (gr *GifReader) GetDecodeImage() *DecodedImage {
	return gr.decodedImg
}

type JpgeReader struct {
	decodedImg *DecodedImage
}

func CreateJpgeReader(image string) *JpgeReader {
	decodedImage := fmt.Sprintf("jpge decoded: %s", image)
	return &JpgeReader{decodedImg: CreateDecodedImage(decodedImage)}
}

func (jr *JpgeReader) GetDecodeImage() *DecodedImage {
	return jr.decodedImg
}

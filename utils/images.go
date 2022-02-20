package utils

import (
	"bytes"
	"image"
	"image/jpeg"
)

func EncodeImage(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}

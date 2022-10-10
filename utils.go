package main

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// https://stackoverflow.com/questions/49594259/reading-image-in-go
func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func getEbitenImage(filepath string) (*ebiten.Image, error) {
	img, err := getImageFromFilePath("tex.png")
	if err != nil {
		return &ebiten.Image{}, err
	}
	return ebiten.NewImageFromImage(img), nil
}

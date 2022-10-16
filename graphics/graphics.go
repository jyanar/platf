package graphics

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var Atlas *ebiten.Image

// In the atlas, numbers are left -> right
var Tile 		*ebiten.Image // 0
var ToggleFloor *ebiten.Image // 1
var Empty 		*ebiten.Image // 2
var Symbol 		*ebiten.Image // 3
var Spikes 		*ebiten.Image // 4
var LeverOff 	*ebiten.Image // 5
var LeverOn 	*ebiten.Image // 6
var Player 		*ebiten.Image // 8
var Enemy 		*ebiten.Image // 20

func Load() error {
	var err error

	// load the texture atlas
	Atlas, err = getEbitenImage("tex.png")
	if err != nil {
		return err
	}

	// pull subimages for all tiles out of the atlas
	Tile, err = getEbitenSubImageAt(Atlas, 0, 0, 16, 16)
	if err != nil { return err }

	ToggleFloor, err = getEbitenSubImageAt(Atlas, 16, 0, 16, 16)
	if err != nil { return err }

	Empty, err = getEbitenSubImageAt(Atlas, 32, 0, 16, 16)
	if err != nil { return err }

	Symbol, err = getEbitenSubImageAt(Atlas, 48, 0, 16, 16)
	if err != nil { return err }

	Spikes, err = getEbitenSubImageAt(Atlas, 0, 16, 16, 16)
	if err != nil { return err }

	LeverOff, err = getEbitenSubImageAt(Atlas, 16, 16, 16, 16)
	if err != nil { return err }

	LeverOn, err = getEbitenSubImageAt(Atlas, 32, 16, 16, 16)
	if err != nil { return err }

	Player, err = getEbitenSubImageAt(Atlas, 32, 32, 16, 16)
	if err != nil { return err }

	Enemy, err = getEbitenSubImageAt(Atlas, 80, 0, 16, 16)
	if err != nil { return err }

	return nil
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

func getEbitenSubImageAt(atlas *ebiten.Image, x, y, width, height int) (*ebiten.Image, error) {
	return atlas.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image), nil
}
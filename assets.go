package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
	tex *ebiten.Image
	// quads []*image.Image
	// quads []*ebiten.Image
}

func (a Assets) qdraw(screen *ebiten.Image, id int, x float64, y float64) {
	tex, err := getEbitenImage("tex.png")
	if err != nil {
		fmt.Println("ERROR!!! ")
		fmt.Println(err)
	}
	// Grab image id from tex
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	texWidth, texHeight := tex.Size()
	// id := 0
	ctr := 0
	for x := 0; x < texWidth; x = x + 16 {
		for y := 0; y < texHeight; y = y + 16 {
			ctr++
			if ctr == id {
				screen.DrawImage(tex.SubImage(image.Rect(x, y, x+16, y+16)).(*ebiten.Image), op)
			}
		}
	}

	// for y := 0; y < texHeight; y = y + 16 {
	// 	for x := 0; x < texWidth; x = x + 16 {
	// 		ctr++
	// 		if ctr == id {
	// 			screen.DrawImage(tex.SubImage(image.Rect(x, y, 16, 16)).(*ebiten.Image), op)
	// 		}
	// 		// img := tex.SubImage(image.Rect(x, y, 16, 16)).(*ebiten.Image)
	// 		// a.quads = append(a.quads, img)
	// 	}
	// }
	// screen.DrawImage(a.quads[id], op)
	// screen.DrawImage(s.ebitenImage.SubImage(image.Rect(0, 32, 16, 16)).(*ebiten.Image), op)
}

func (a *Assets) init() error {
	t, err := getEbitenImage("tex.png")
	if err != nil {
		fmt.Println(err)
	}
	a.tex = t
	return nil
}

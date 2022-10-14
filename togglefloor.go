package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ToggleFloor struct {
	Obj
	isSolid bool
	assets  Assets
}

func NewToggleFloor(obj Obj) *ToggleFloor {
	return &ToggleFloor{obj, true, Assets{}}
}

func (t ToggleFloor) Solid() bool {
	return t.isSolid
}

func (t ToggleFloor) Update(state *GameState) error { return nil }

func (t ToggleFloor) Draw(screen *ebiten.Image) {
	if t.Solid() {
		t.assets.qdraw(screen, 7, t.x, t.y)
	} else {
		ebitenutil.DrawRect(screen, t.x, t.y, 3, 3, image.White)
	}
}

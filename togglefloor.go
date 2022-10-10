package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ToggleFloor struct {
	Obj
	isSolid bool
}

func NewToggleFloor(obj Obj) *ToggleFloor {
	return &ToggleFloor{obj, true}
}

func (t ToggleFloor) Solid() bool {
	return t.isSolid
}

func (t ToggleFloor) Update(state *GameState) error { return nil }

func (t ToggleFloor) Draw(screen *ebiten.Image) {
	if t.Solid() {
		ebitenutil.DrawCircle(screen, t.x+t.w/2, t.y+t.h/2, t.w/2, image.White)
	} else {
		ebitenutil.DrawRect(screen, t.x, t.y, 3, 3, image.White)
	}
}

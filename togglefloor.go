package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.x, t.y)
	if t.Solid() {
		screen.DrawImage(graphics.ToggleFloor, op)
	} else {
		screen.DrawImage(graphics.Empty, op)
	}
}

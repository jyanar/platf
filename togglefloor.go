package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type ToggleFloor struct {
	Object
}

func NewToggleFloor(Object Object) *ToggleFloor {
	Object.isSolid = false
	return &ToggleFloor{Object}
}

func (t ToggleFloor) Solid() bool {
	return t.Object.Solid()
}

func (t ToggleFloor) Update(state *GameState) error { return nil }

func (t ToggleFloor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.position.x, t.position.y)
	if t.Solid() {
		screen.DrawImage(graphics.Quads[1], op)
	} else {
		screen.DrawImage(graphics.Quads[2], op)
	}
}

func (t *ToggleFloor) toggleSolid() {
	t.isSolid = !t.isSolid
}

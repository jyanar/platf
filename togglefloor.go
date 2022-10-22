package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type ToggleFloor struct {
	Obj
}

func NewToggleFloor(obj Obj) *ToggleFloor {
	return &ToggleFloor{obj}
}

func (t ToggleFloor) Solid() bool {
	return t.Obj.Solid()
}

func (t ToggleFloor) Update(state *GameState) error { return nil }

func (t ToggleFloor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.x, t.y)
	if t.Solid() {
		screen.DrawImage(graphics.Quads[1], op)
	} else {
		screen.DrawImage(graphics.Quads[2], op)
	}
}

func (t *ToggleFloor) toggleSolid() {
	t.isSolid = !t.isSolid
}

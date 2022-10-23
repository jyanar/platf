package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Portal struct {
	Obj
}

func NewPortal(obj Obj) *Portal {
	return &Portal{obj}
}

func (p Portal) Update(state *GameState) error { return nil }

func (p Portal) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(graphics.Portal, op)
}

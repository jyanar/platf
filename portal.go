package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Portal struct {
	Object
}

func NewPortal(Object Object) *Portal {
	return &Portal{Object}
}

func (p Portal) Update(state *GameState) error { return nil }

func (p Portal) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.x, p.position.y)
	screen.DrawImage(graphics.Portal, op)
}

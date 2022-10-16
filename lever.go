package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Lever struct {
	Obj
	toggle  bool
	isSolid bool
}

func NewLever(obj Obj) *Lever {
	return &Lever{obj, false, false}
}

func (l Lever) Solid() bool {
	return l.isSolid
}

func (l Lever) Update(state *GameState) error { return nil }

func (l Lever) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.x, l.y)
	screen.DrawImage(graphics.LeverOn, op)
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Lever struct {
	Object
	toggle bool
}

func NewLever(Object Object) *Lever {
	return &Lever{Object, false}
}

func (l Lever) Solid() bool {
	return l.Object.Solid()
}

func (l Lever) Update(state *GameState) error { return nil }

func (l Lever) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.position.x, l.position.y)
	if l.toggle {
		screen.DrawImage(graphics.LeverOn, op)
	} else {
		screen.DrawImage(graphics.LeverOff, op)
	}
}

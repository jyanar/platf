package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ebitenutil.DrawCircle(screen, l.x+l.w/2, l.y+l.h/2, l.w/2, image.White)
}
